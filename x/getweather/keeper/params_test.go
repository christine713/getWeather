package keeper_test

import (
	"testing"

	testkeeper "github.com/christine713/getWeather/testutil/keeper"
	"github.com/christine713/getWeather/x/getweather/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.GetweatherKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
