package keeper

import (
	"context"

	"github.com/christine713/getWeather/x/getweather/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) QueryWether(goCtx context.Context, msg *types.MsgQueryWether) (*types.MsgQueryWetherResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgQueryWetherResponse{}, nil
}
