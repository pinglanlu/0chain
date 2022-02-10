package stakepool

import (
	"errors"
	"fmt"

	"0chain.net/core/util"

	"0chain.net/chaincore/state"

	cstate "0chain.net/chaincore/chain/state"
	"0chain.net/chaincore/transaction"
	"0chain.net/core/datastore"
)

func CheckClientBalance(
	t *transaction.Transaction,
	balances cstate.StateContextI,
) (err error) {
	if t.Value < 0 {
		return errors.New("negative transaction value")
	}

	var balance state.Balance
	balance, err = balances.GetClientBalance(t.ClientID)

	if err != nil && err != util.ErrValueNotPresent {
		return
	}

	if err == util.ErrValueNotPresent {
		return errors.New("no tokens to lock")
	}

	if state.Balance(t.Value) > balance {
		return errors.New("lock amount is greater than balance")
	}

	return
}

func (sp *StakePool) LockPool(
	txn *transaction.Transaction,
	providerType Provider,
	providerId datastore.Key,
	status PoolStatus,
	balances cstate.StateContextI,
) error {
	const MaxDelegates = 100

	if err := CheckClientBalance(txn, balances); err != nil {
		return err
	}

	if len(sp.Pools) >= MaxDelegates {
		return fmt.Errorf("max_delegates reached: %v, no more stake pools allowed",
			MaxDelegates)
	}

	dp := DelegatePool{
		Balance:      state.Balance(txn.Value),
		Reward:       0,
		Status:       status,
		DelegateID:   txn.ClientID,
		RoundCreated: balances.GetBlock().Round,
	}

	if err := balances.AddTransfer(state.NewTransfer(
		txn.ClientID, txn.ToClientID, state.Balance(txn.Value),
	)); err != nil {
		return err
	}

	var newPoolId = txn.Hash
	sp.Pools[newPoolId] = &dp

	var usp *UserStakePools
	usp, err := getOrCreateUserStakePool(providerType, txn.ClientID, balances)
	if err != nil {
		return fmt.Errorf("can't get user pools list: %v", err)
	}
	usp.add(providerId, newPoolId)
	if err = usp.Save(providerType, txn.ClientID, balances); err != nil {
		return fmt.Errorf("saving user pools: %v", err)
	}

	if err := dp.emitNew(
		txn.ClientID,
		newPoolId,
		providerId,
		providerType,
		status,
		balances,
	); err != nil {
		return err
	}

	return nil
}
