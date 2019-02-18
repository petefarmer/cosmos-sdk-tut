package nameservice

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgSetName struct {
	Name string
	Value string
	Owner sdk.AccAddress
}

type MsgBuyName struct {
	Name string
	Bid	sdk.Coins
	Buyer	sdk.AccAddress
}

func NewMsgBuyName(name string, bid sdk.Coins, buyer sdk.AccAddress) MsgBuyName {
	return MsgBuyName{
		Name: name,
		Bid: bid,
		Buyer: buyer,
	}
}
func NewMsgSetName(name string, value string, owner sdk.AccAddress) MsgSetName {
	return MsgSetName{
		Name: name,
		Value: value,
		Owner: owner,
	}
}

func (msg MsgSetName) Route() string { return "nameservice" }

func (msg MsgSetName) Type() string {return "set_name"}

func (msg MsgSetName) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Name) == 0 || len(msg.Value) == 0 {
		return sdk.ErrUnknownREquest("Name and/or Value cannot be empty")
	}
	return nil
}

func (msg MsgSetName) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}

