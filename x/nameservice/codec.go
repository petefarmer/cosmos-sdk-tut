package nameservice

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSetNme{}, "nameservice/SetName", nil)
	cdc.RegisterConcrete(MsgBuyNme{}, "nameservice/BuyName", nil)
}
