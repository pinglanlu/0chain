package storagesc

import (
	"fmt"

	state "0chain.net/chaincore/chain/state"
	"0chain.net/smartcontract/partitions"
	"github.com/0chain/common/core/logging"
	"go.uber.org/zap"
)

//go:generate msgp -io=false -tests=false -v

//------------------------------------------------------------------------------

// BlobberAllocationNode represents the allocation that belongs to a blobber,
// will be saved in blobber allocations partitions.
type BlobberAllocationNode struct {
	ID string `json:"id"` // allocation id
}

func (z *BlobberAllocationNode) GetID() string {
	return z.ID
}

func partitionsBlobberAllocations(blobberID string, balances state.StateContextI) (*partitions.Partitions, error) {
	return partitions.CreateIfNotExists(balances, getBlobberAllocationsKey(blobberID), blobberAllocationPartitionSize)
}

func partitionsBlobberAllocationsAdd(state state.StateContextI, blobberID, allocID string) error {
	blobAllocsParts, err := partitionsBlobberAllocations(blobberID, state)
	if err != nil {
		return fmt.Errorf("error fetching blobber challenge allocation partition, %v", err)
	}

	err = blobAllocsParts.Add(state, &BlobberAllocationNode{ID: allocID})
	if err != nil && !partitions.ErrItemExist(err) {
		return err
	} else if partitions.ErrItemExist(err) {
		return nil
	}

	if err := blobAllocsParts.Save(state); err != nil {
		return fmt.Errorf("could not update blobber allocations partitions: %v", err)
	}

	return nil
}

// removeAllocationFromBlobberPartitions removes the allocation from blobber
func removeAllocationFromBlobberPartitions(state state.StateContextI, blobberID, allocID string) error {
	logging.Logger.Info("remove allocation from blobber partitions",
		zap.String("blobber", blobberID),
		zap.String("allocation", allocID))

	blobAllocsParts, err := partitionsBlobberAllocations(blobberID, state)
	if err != nil {
		return fmt.Errorf("could not get blobber allocations partition: %v", err)
	}

	err = blobAllocsParts.Remove(state, allocID)

	if err == nil {
		logging.Logger.Info("error is not there in remove allocation from blobber partitions",
			zap.String("blobber", blobberID),
			zap.String("allocation", allocID))

		if err := blobAllocsParts.Save(state); err != nil {
			logging.Logger.Info("could not update blobber allocation partitions",
				zap.Error(err),
				zap.String("blobber", blobberID),
				zap.String("allocation", allocID))
			return fmt.Errorf("could not update blobber allocation partitions: %v", err)
		}

		baParts, err := partitionsBlobberAllocations(blobberID, state)
		logging.Logger.Info("blobber allocation partition size",
			zap.Any("baParts", baParts),
			zap.String("blobber", blobberID),
			zap.String("allocation", allocID))
		if err != nil {
			return fmt.Errorf("could not get blobber allocations partition: %v", err)
		}

		allocNum, err := blobAllocsParts.Size(state)
		if err != nil {
			return fmt.Errorf("could not get challenge partition size: %v", err)
		}

		logging.Logger.Info("blobber allocation partition size",
			zap.Int("size", allocNum),
			zap.String("blobber", blobberID),
			zap.String("allocation", allocID))

		if allocNum > 0 {
			return nil
		}

		// remove blobber from challenge ready partition when there's no allocation bind to it
		err = partitionsChallengeReadyBlobbersRemove(state, blobberID)
		if err != nil && !partitions.ErrItemNotFound(err) {
			// it could be empty if we finalize the allocation before committing any read or write
			return fmt.Errorf("failed to remove blobber from challenge ready partitions: %v", err)
		}

		return nil
	} else {
		if !partitions.ErrItemNotFound(err) {
			logging.Logger.Error("could not remove allocation from blobber",
				zap.Error(err),
				zap.String("blobber", blobberID),
				zap.String("allocation", allocID))
		} else {
			logging.Logger.Error("allocation is not in partition",
				zap.Error(err),
				zap.String("blobber", blobberID),
				zap.String("allocation", allocID))
		}
		return fmt.Errorf("could not remove allocation from blobber: %v", err)
	}
}
