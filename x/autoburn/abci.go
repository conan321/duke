package autoburn

import (
	"duke/x/autoburn/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func EndBlock(ctx sdk.Context, k keeper.Keeper) {
	if err := k.MaybeBurnCoins(ctx); err != nil {
		k.Logger(ctx).Error("auto burn coins error", "err", err)
	}
}
