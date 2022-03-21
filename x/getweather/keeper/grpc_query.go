package keeper

import (
	"github.com/christine713/getWeather/x/getweather/types"
)

var _ types.QueryServer = Keeper{}
