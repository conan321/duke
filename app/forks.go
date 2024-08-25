package app

import (
	v11 "duke/app/upgrades/v11"
	autoburntypes "duke/x/autoburn/types"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

// BeginBlockForks executes any necessary fork logic based upon the current block height.
func BeginBlockForks(ctx sdk.Context, app *App) {
	if ctx.BlockHeight() == v11.HaltBlockHeight {
		v11.UpdateGovMinDepositParam(ctx, app.GovKeeper)

		storeUpgrades := storetypes.StoreUpgrades{
			Added: []string{autoburntypes.ModuleName},
		}

		// configure store loader that checks if version == upgradeHeight and applies store upgrades
		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(v11.HaltBlockHeight, &storeUpgrades))
	}
}
