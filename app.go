package app

import (
	"github.com/tendermint/tendermint/libs/log"
	"github.com/cosmos/cosmos-sdk/x/auth"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	dbm "github.com/tendermint/tendermint/libs/db"
)

const (
	appName = "nameservice"
)

type nameServiceApp struct {
	*bam.BaseApp
}

func NewNameServiceApp(logger log.Logger, db dbm.DB) *nameServiceApp {
  // define top-level shared codec
  cdc := MakeCodec()

  // BaseApp handler
  bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTXDecoder(cdc))

  var app = &nameServiceApp{
    BaseApp: bApp,
    cdc:     cdc,
  }

  return app
}
