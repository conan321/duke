package v11

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
)

func UpdateGovMinDepositParam(ctx sdk.Context, keeper govkeeper.Keeper) {
	params := keeper.GetParams(ctx)
	coins := []sdk.Coin{}
	for _, minDeposit := range params.MinDeposit {
		coins = append(coins, sdk.NewCoin(MinDepositDenom, minDeposit.Amount))
	}
	params.MinDeposit = coins
	keeper.SetParams(ctx, params)
}
