package keeper

import (
	"github.com/code0xff/planet/x/planet/types"
)

var _ types.QueryServer = Keeper{}
