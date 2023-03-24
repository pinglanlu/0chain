package event

import (
	"0chain.net/smartcontract/common"
	"0chain.net/smartcontract/dbs"
	"0chain.net/smartcontract/dbs/model"
	"0chain.net/smartcontract/stakepool/spenum"
	"github.com/0chain/common/core/currency"
	"github.com/0chain/common/core/logging"
	"go.uber.org/zap"
	"gorm.io/gorm/clause"
)

// swagger:model RewardProvider
type RewardProvider struct {
	model.UpdatableModel
	Amount      currency.Coin `json:"amount"`
	BlockNumber int64         `json:"block_number" gorm:"index:idx_rew_block_prov,priority:1"`
	ProviderId  string        `json:"provider_id" gorm:"index:idx_rew_block_prov,priority:2"`
	RewardType  spenum.Reward `json:"reward_type"`
	ChallengeID string        `json:"challenge_id"`
}

func (edb *EventDb) insertProviderReward(inserts []dbs.StakePoolReward, round int64) error {
	if len(inserts) == 0 {
		return nil
	}

	logging.Logger.Debug("insertProviderReward", zap.Any("inserts", inserts), zap.Any("round", round))

	var prs []RewardProvider
	for _, sp := range inserts {
		pr := RewardProvider{
			Amount:      sp.Reward,
			BlockNumber: round,
			ProviderId:  sp.ProviderId,
			RewardType:  sp.RewardType,
			ChallengeID: sp.ChallengeID,
		}

		prs = append(prs, pr)
	}
	err := edb.Get().Create(&prs).Error

	logging.Logger.Debug("insertProviderReward", zap.Any("err", err))

	return err
}

func (edb *EventDb) GetProviderRewards(limit common.Pagination, id string, start, end int64) ([]RewardProvider, error) {
	var rps []RewardProvider
	query := edb.Get().Model(&RewardProvider{})
	if id == "" {
		if start == end {
			query = query.Where("block_number = ?", start)
		} else {
			query = query.Where("block_number >= ? AND block_number < ?", start, end)
		}
	} else {
		if start == end {
			query = query.Where("provider_id = ? AND block_number = ?", id, start)
		} else {
			query = query.Where("provider_id = ? AND block_number >= ? AND block_number < ?", id, start, end)
		}
	}

	return rps, query.Offset(limit.Offset).
		Limit(limit.Limit).
		Order(clause.OrderByColumn{
			Column: clause.Column{Name: "block_number"},
			Desc:   limit.IsDescending,
		}).Scan(&rps).Error
}

func (edb *EventDb) GetChallengeRewardsByChallengeID(challengeID string) []RewardProvider {

	var rps []RewardProvider
	edb.Get().Where("challenge_id = ? AND reward_type IN (6, 8, 9)", challengeID).Find(&rps)

	return rps
}

func (edb *EventDb) GetChallengeRewardsByProviderID(providerID string) []RewardProvider {

	var rps []RewardProvider
	edb.Get().Where("provider_id = ? AND reward_type IN (0, 6, 8, 9, 10)", providerID).Find(&rps)

	return rps
}

func (edb *EventDb) GetAllProviderChallengeRewards() []RewardProvider {

	var rps []RewardProvider
	edb.Get().Where("reward_type IN (0, 6, 8, 9, 10)").Find(&rps)

	return rps
}

func (edb *EventDb) GetRewardsByRewardType(rewardType string) []RewardProvider {

	var rps []RewardProvider
	edb.Get().Where("reward_type = ?", rewardType).Find(&rps)

	return rps
}

func (edb *EventDb) RunWhereQueryInProviderRewards(query string) []RewardProvider {

	var rps []RewardProvider

	edb.Get().Where(query).Find(&rps)

	return rps

}

func (edb *EventDb) GetBlockRewardsToProviders(blockNumber, startBlockNumber, endBlockNumber string) []RewardProvider {

	if blockNumber != "" {
		var rps []RewardProvider
		edb.Get().Where("block_number = ? AND reward_type IN (3)", blockNumber).Find(&rps)

		return rps
	}

	if startBlockNumber != "" && endBlockNumber != "" {
		var rps []RewardProvider
		edb.Get().Where("block_number >= ? AND block_number <= ? AND reward_type IN (3)", startBlockNumber, endBlockNumber).Find(&rps)

		return rps
	}

	return nil
}

func (edb *EventDb) GetReadRewardsToProviders(blockNumber, startBlockNumber, endBlockNumber string) []RewardProvider {

	if blockNumber != "" {
		var rps []RewardProvider
		edb.Get().Where("block_number = ? AND reward_type IN (7)", blockNumber).Find(&rps)

		return rps
	}

	if startBlockNumber != "" && endBlockNumber != "" {
		var rps []RewardProvider
		edb.Get().Where("block_number >= ? AND block_number <= ? AND reward_type IN (7)", startBlockNumber, endBlockNumber).Find(&rps)

		return rps
	}

	return nil
}
