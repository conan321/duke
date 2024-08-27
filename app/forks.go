package app

import (
	v11 "duke/app/upgrades/v11"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BeginBlockForks executes any necessary fork logic based upon the current block height.
func BeginBlockForks(ctx sdk.Context, app *App) {
	if ctx.BlockHeight() == v11.ForkBlockHeight {
		v11.UpdateGovMinDepositParam(ctx, app.GovKeeper)
	}
}
