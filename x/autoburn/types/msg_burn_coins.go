package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgBurnCoins = "burn_coins"
)

var (
	_ sdk.Msg = &MsgBurnCoins{}
)

func NewMsgBurnCoins(sender string, coins []sdk.Coin) *MsgBurnCoins {
	return &MsgBurnCoins{
		Sender: sender,
		Coins:  coins,
	}
}

func (m *MsgBurnCoins) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(m.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (m *MsgBurnCoins) ValidateBasic() error {
	if len(m.Coins) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "invalid coins")
	}
	return nil
}
