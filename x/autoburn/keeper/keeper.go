package keeper

import (
	"fmt"

	"github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	"duke/x/autoburn/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   storetypes.StoreKey
		memKey     storetypes.StoreKey
		paramstore paramtypes.Subspace

		ak types.AccountKeeper
		bk types.BankKeeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey storetypes.StoreKey,
	ps paramtypes.Subspace,
	ak types.AccountKeeper,
	bk types.BankKeeper,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,

		ak: ak,
		bk: bk,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) BurnCoins(ctx sdk.Context, sender sdk.AccAddress, amounts sdk.Coins) error {
	return k.bk.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, amounts)
}

func (k Keeper) MaybeBurnCoins(ctx sdk.Context) error {
	// Get module account balance
	moduleAccount := k.ak.GetModuleAccount(ctx, types.ModuleName)
	balances := k.bk.GetAllBalances(ctx, moduleAccount.GetAddress())
	if balances.IsZero() {
		k.Logger(ctx).Info("no balances for burn")
		return nil
	}
	params := k.GetParams(ctx)
	if len(params.Thresholds) == 0 {
		k.Logger(ctx).Info("burn all", "coins", balances.String())
		return k.bk.BurnCoins(ctx, types.ModuleName, balances)
	}
	return k.burnWithThreshold(ctx, balances, params.Thresholds)
}

func (k Keeper) burnWithThreshold(ctx sdk.Context, balances sdk.Coins, thresholds sdk.Coins) error {
	coins := sdk.Coins{}
	for _, balance := range balances {
		if balance.Amount.IsPositive() && IsGTE(balance, thresholds) {
			coins.Add(balance)
		}
	}
	if !coins.IsZero() {
		if err := k.bk.BurnCoins(ctx, types.ModuleName, coins); err != nil {
			return err
		}
	}
	return nil
}

func IsGTE(balance sdk.Coin, thresholds sdk.Coins) bool {
	for _, threshold := range thresholds {
		if balance.Denom == threshold.Denom && balance.Amount.GTE(threshold.Amount) {
			return true
		}
	}
	return false
}
