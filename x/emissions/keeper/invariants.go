package keeper

import (
	"fmt"

	cosmosMath "cosmossdk.io/math"
	"github.com/allora-network/allora-chain/app/params"
	emissionsTypes "github.com/allora-network/allora-chain/x/emissions/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// RegisterInvariants registers the emissions module invariants.
func RegisterInvariants(ir sdk.InvariantRegistry, k *Keeper) {
	ir.RegisterRoute(emissionsTypes.ModuleName, "allora-staking-total-balance", StakingInvariantTotalStakeEqualAlloraStakingBankBalance(*k))
	ir.RegisterRoute(emissionsTypes.ModuleName, "allora-staking-total-stake-equal-topic-stake", StakingInvariantTotalStakeEqualSumTopicStakes(*k))
	ir.RegisterRoute(emissionsTypes.ModuleName, "stake-removals-length-same", StakingInvariantLenStakeRemovalsSame(*k))
}

// AllInvariants is a convience function to run all invariants in the emissions module.
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		if res, stop := StakingInvariantTotalStakeEqualAlloraStakingBankBalance(k)(ctx); stop {
			return res, stop
		}
		if res, stop := StakingInvariantTotalStakeEqualSumTopicStakes(k)(ctx); stop {
			return res, stop
		}
		if res, stop := StakingInvariantLenStakeRemovalsSame(k)(ctx); stop {
			return res, stop
		}
		return "", false
	}
}

// StakingInvariantTotalStakeEqualAlloraStakingBankBalance checks that
// the total stake in the emissions module is equal to the balance
// of the Allora staking bank account.
func StakingInvariantTotalStakeEqualAlloraStakingBankBalance(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		totalStake, err := k.GetTotalStake(ctx)
		if err != nil {
			panic(fmt.Sprintf("failed to get total stake: %v", err))
		}
		alloraStakingAddr := k.authKeeper.GetModuleAccount(ctx, emissionsTypes.AlloraStakingAccountName).GetAddress()
		alloraStakingBalance := k.bankKeeper.GetBalance(
			ctx,
			alloraStakingAddr,
			params.DefaultBondDenom).Amount
		broken := !totalStake.Equal(alloraStakingBalance)
		return sdk.FormatInvariant(
			emissionsTypes.ModuleName,
			"emissions module total stake equal allora staking bank balance",
			fmt.Sprintf("TotalStake: %s | Allora Module Account Staking Balance: %s",
				totalStake.String(),
				alloraStakingBalance.String(),
			),
		), broken
	}
}

// the sum of all topicStake should always be equal to totalStake
func StakingInvariantTotalStakeEqualSumTopicStakes(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		totalStake, err := k.GetTotalStake(ctx)
		if err != nil {
			panic(fmt.Sprintf("failed to get total stake: %v", err))
		}
		numTopics, err := k.GetNextTopicId(ctx)
		if err != nil {
			panic(fmt.Sprintf("failed to get next topic id: %v", err))
		}
		sumTopicStakes := cosmosMath.ZeroInt()
		for i := uint64(0); i < numTopics; i++ {
			topicStake, err := k.GetTopicStake(ctx, i)
			if err != nil {
				panic(fmt.Sprintf("failed to get topic stake: %v", err))
			}
			sumTopicStakes = sumTopicStakes.Add(topicStake)
		}
		broken := !totalStake.Equal(sumTopicStakes)
		return sdk.FormatInvariant(
			emissionsTypes.ModuleName,
			"emissions module total stake equal sum of all topic stakes",
			fmt.Sprintf("TotalStake: %s | Sum of all Topic Stakes: %s | num topics :%d",
				totalStake.String(),
				sumTopicStakes.String(),
				numTopics,
			),
		), broken
	}
}

// the number of values in the stakeRemovalsByBlock map
// should always equal the number of values in the stakeRemovalsByActor map
func StakingInvariantLenStakeRemovalsSame(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		iterByBlock, err := k.stakeRemovalsByBlock.Iterate(ctx, nil)
		if err != nil {
			panic(fmt.Sprintf("failed to get stake removals iterator: %v", err))
		}
		valuesByBlock, err := iterByBlock.Values()
		if err != nil {
			panic(fmt.Sprintf("failed to get stake removals values: %v", err))
		}
		lenByBlock := len(valuesByBlock)
		iterByActor, err := k.stakeRemovalsByActor.Iterate(ctx, nil)
		if err != nil {
			panic(fmt.Sprintf("failed to get stake removals iterator: %v", err))
		}
		valuesByActor, err := iterByActor.Keys()
		if err != nil {
			panic(fmt.Sprintf("failed to get stake removals values: %v", err))
		}
		lenByActor := len(valuesByActor)

		broken := lenByBlock != lenByActor
		return sdk.FormatInvariant(
			emissionsTypes.ModuleName,
			"emissions module length of stake removals same",
			fmt.Sprintf("Length of stake removals: %d | Length of stake removals: %d\n%v\n%v",
				lenByBlock,
				lenByActor,
				valuesByBlock,
				valuesByActor,
			),
		), broken
	}
}
