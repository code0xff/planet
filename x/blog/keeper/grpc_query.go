package keeper

import (
	"github.com/code0xff/planet/x/blog/types"
)

var _ types.QueryServer = Keeper{}
