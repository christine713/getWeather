package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgQueryWether = "query_wether"

var _ sdk.Msg = &MsgQueryWether{}

func NewMsgQueryWether(creator string, title string, body string) *MsgQueryWether {
	return &MsgQueryWether{
		Creator: creator,
		Title:   title,
		Body:    body,
	}
}

func (msg *MsgQueryWether) Route() string {
	return RouterKey
}

func (msg *MsgQueryWether) Type() string {
	return TypeMsgQueryWether
}

func (msg *MsgQueryWether) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgQueryWether) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgQueryWether) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
