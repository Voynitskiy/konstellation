// nolint
package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type CodeType = sdk.CodeType

const (
	DefaultCodespace            sdk.CodespaceType = "issue"
	CodeInvalidGenesis          sdk.CodeType      = 102
	CodeUnknownIssue            sdk.CodeType      = 1
	CodeIssuerMismatch          sdk.CodeType      = 2
	CodeInvalidDenom            sdk.CodeType      = 3
	CodeAmountLowerAllowance    sdk.CodeType      = 4
	CodeIssueExists             sdk.CodeType      = 5
	CodeNotEnoughFee            sdk.CodeType      = 6
	CodeInvalidFeature          sdk.CodeType      = 7
	CodeCanNotMint              sdk.CodeType      = 8
	CodeCanNotBurnOwner         sdk.CodeType      = 9
	CodeCanNotBurnHolder        sdk.CodeType      = 10
	CodeCanNotBurnFrom          sdk.CodeType      = 11
	CodeCanNotFreeze            sdk.CodeType      = 12
	CodeAmountNotValid          sdk.CodeType      = 13
	CodeInvalidCoinDecimals     sdk.CodeType      = 14
	CodeInvalidTotalSupply      sdk.CodeType      = 15
	CodeInvalidDescription      sdk.CodeType      = 16
	CodeInvalidSymbol           sdk.CodeType      = 17
	CodeInvalidFreezeOp         sdk.CodeType      = 18
	CodeNotTransferOut          sdk.CodeType      = 19
	CodeNotTransferIn           sdk.CodeType      = 20
	CodeInvalidInput            sdk.CodeType      = 400
	CodeInvalidIssueFee         sdk.CodeType      = 401
	CodeInvalidMintFee          sdk.CodeType      = 402
	CodeInvalidBurnFee          sdk.CodeType      = 402
	CodeInvalidBurnFromFee      sdk.CodeType      = 403
	CodeInvalidFreezeFee        sdk.CodeType      = 404
	CodeInvalidUnFreezeFee      sdk.CodeType      = 405
	CodeInvalidTransferOwnerFee sdk.CodeType      = 406
)

var (
	ErrInvalidIssueParams               = sdkerrors.Register(DefaultCodespace, CodeInvalidInput, "Invalid issue params")
	ErrIssueAlreadyExists               = sdkerrors.Register(DefaultCodespace, CodeIssueExists, "Invalid already exists")
	ErrNotEnoughFee                     = sdkerrors.Register(DefaultCodespace, CodeNotEnoughFee, "Not enough fee")
	ErrCoinDecimalsMaxValueNotValid     = sdkerrors.Register(DefaultCodespace, CodeInvalidCoinDecimals, fmt.Sprintf("Decimals max value is %d", CoinDecimalsMaxValue))
	ErrCoinDecimalsMultipleNotValid     = sdkerrors.Register(DefaultCodespace, CodeInvalidCoinDecimals, fmt.Sprintf("Decimals must be a multiple of %d", CoinDecimalsMultiple))
	ErrCoinTotalSupplyMaxValueNotValid  = sdkerrors.Register(DefaultCodespace, CodeInvalidTotalSupply, fmt.Sprintf("Total supply max value is %s", CoinMaxTotalSupply.String()))
	ErrCoinDescriptionNotValid          = sdkerrors.Register(DefaultCodespace, CodeInvalidDescription, "Description is not valid json")
	ErrCoinDescriptionMaxLengthNotValid = sdkerrors.Register(DefaultCodespace, CodeInvalidDescription, fmt.Sprintf("Description max length is %d", CoinDescriptionMaxLength))
	ErrCoinSymbolNotValid               = sdkerrors.Register(DefaultCodespace, CodeInvalidSymbol, "Invalid symbol")
)

//convert sdk.Error to error
func Errorf(err sdk.Error) error {
	return fmt.Errorf(err.Stacktrace().Error())
}

func ErrUnknownIssue(denom string) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeUnknownIssue, fmt.Sprintf("Unknown issue %s", denom))
}

func ErrOwnerMismatch(issueID string) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeIssuerMismatch, fmt.Sprintf("Owner mismatch with token %s", issueID))
}

func ErrAmountGreaterThanAllowance(amt sdk.Coin, allowance sdk.Coin) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeAmountLowerAllowance, fmt.Sprintf("Amount greater than allowance %s > %s", amt.String(), allowance.String()))
}

func ErrInvalidDenom(denom string) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeInvalidDenom, fmt.Sprintf("Denom invalid %s", denom))
}

func ErrInvalidFreezeOp(op string) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeInvalidFreezeOp, fmt.Sprintf("Invalid freeze type %s", op))
}

func ErrInvalidFeature(feature string) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeInvalidFeature, fmt.Sprintf("Feature invalid %s", feature))
}

func ErrCanNotMint(denom string) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeCanNotMint, fmt.Sprintf("Can not mint the token %s", denom))
}

func ErrCanNotBurnOwner(denom string) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeCanNotBurnOwner, fmt.Sprintf("Can not burn the token %s", denom))
}

func ErrCanNotBurnHolder(denom string) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeCanNotBurnHolder, fmt.Sprintf("Can not burn the token %s", denom))
}

func ErrCanNotBurnFrom(denom string) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeCanNotBurnFrom, fmt.Sprintf("Can not burn the token %s", denom))
}

func ErrCanNotFreeze(denom string) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeCanNotFreeze, fmt.Sprintf("Can not freeze the token %s", denom))
}

func ErrCanNotTransferIn(denom string, accAddress string) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeNotTransferIn, fmt.Sprintf("Can not transfer in %s to %s", denom, accAddress))
}

func ErrCanNotTransferOut(denom string, accAddress string) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeNotTransferOut, fmt.Sprintf("Can not transfer out %s from %s", denom, accAddress))
}

func ErrInvalidIssueFee(fee string) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeInvalidIssueFee, fmt.Sprintf("invalid issue fee: %s", fee))
}

func ErrInvalidMintFee(fee string) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeInvalidMintFee, fmt.Sprintf("invalid mint fee: %s", fee))
}

func ErrInvalidBurnFee(fee string) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeInvalidBurnFee, fmt.Sprintf("invalid burn fee: %s", fee))
}

func ErrInvalidBurnFromFee(fee string) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeInvalidBurnFromFee, fmt.Sprintf("invalid burn from fee: %s", fee))
}

func ErrInvalidFreezeFee(fee string) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeInvalidFreezeFee, fmt.Sprintf("invalid freeze fee: %s", fee))
}

func ErrInvalidUnfreezeFee(fee string) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeInvalidUnFreezeFee, fmt.Sprintf("invalid unfreeze fee: %s", fee))
}

func ErrInvalidTransferOwnerFee(fee string) sdk.Error {
	return sdk.NewError(DefaultCodespace, CodeInvalidTransferOwnerFee, fmt.Sprintf("invalid transfer owner fee: %s", fee))
}
