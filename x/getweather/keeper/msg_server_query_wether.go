package keeper

import (
	"context"

	"github.com/christine713/getWeather/x/getweather/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) QueryWether(goCtx context.Context, msg *types.MsgQueryWether) (*types.MsgQueryWetherResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	 // Create variable of type Post
	 var post = types.Post{
		Creator: msg.Creator,
		Title:   msg.Title,
		Body:    msg.Body,
	 }
	 // Add a post to the store and get back the ID
	 id := k.AppendPost(ctx, post)

	return &types.MsgQueryWetherResponse{}, nil
}
