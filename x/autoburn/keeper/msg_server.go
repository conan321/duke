package keeper

import (
	"context"
	"duke/x/autoburn/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	Keeper
}

// BurnCoins implements types.MsgServer.
func (m msgServer) BurnCoins(goCtx context.Context, msg *types.MsgBurnCoins) (*types.MsgBurnCoinsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	// Transfer coins from the user to the module account
	err = m.Keeper.BurnCoins(ctx, sender, msg.Coins)
	if err != nil {
		return nil, err
	}

	return &types.MsgBurnCoinsResponse{}, nil
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}
