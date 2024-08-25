package v11

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
)

func UpdateGovMinDepositParam(ctx sdk.Context, keeper govkeeper.Keeper) {
	params := keeper.GetParams(ctx)
	newMinDeposit := []sdk.Coin{}
	for _, minDeposit := range params.MinDeposit {
		newMinDeposit = append(newMinDeposit, sdk.NewCoin(MinDepositDenom, minDeposit.Amount))
	}
	params.MinDeposit = newMinDeposit
	keeper.SetParams(ctx, params)
}
