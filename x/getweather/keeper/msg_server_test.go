package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/christine713/getWeather/testutil/keeper"
	"github.com/christine713/getWeather/x/getweather/keeper"
	"github.com/christine713/getWeather/x/getweather/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.GetweatherKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
