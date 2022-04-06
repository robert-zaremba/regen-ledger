package utils

import (
	"context"

	"github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"

	"github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

// GetCreditTypeFromBatchDenom extracts the classId from a batch denom string, then retrieves it from the params.
func GetCreditTypeFromBatchDenom(ctx context.Context, store ecocreditv1.StateStore, k ecocredit.ParamKeeper, denom string, creditTypes []*core.CreditType) (core.CreditType, error) {
	sdkCtx := types.UnwrapSDKContext(ctx)
	classId := ecocredit.GetClassIdFromBatchDenom(denom)
	classInfo, err := store.ClassInfoTable().GetByName(ctx, classId)
	if err != nil {
		return core.CreditType{}, err
	}
	if len(creditTypes) == 0 {
		k.Get(sdkCtx, core.KeyCreditTypes, creditTypes)
	}
	return GetCreditType(classInfo.CreditType, creditTypes)
}

// GetCreditType searches for a credit type that matches the given abbreviation within a credit type slice.
func GetCreditType(ctAbbrev string, creditTypes []*core.CreditType) (core.CreditType, error) {
	//creditTypeName = ecocredit.NormalizeCreditTypeName(creditTypeName)
	for _, creditType := range creditTypes {
		// credit type name's stored via params have enforcement on normalization, so we can be sure they will already
		// be normalized here.
		if creditType.Abbreviation == ctAbbrev {
			return *creditType, nil
		}
	}
	return core.CreditType{}, errors.ErrInvalidType.Wrapf("%s is not a valid credit type", ctAbbrev)
}

// GetNonNegativeFixedDecs takes an arbitrary amount of decimal strings, and returns their corresponding fixed decimals
// in a slice.
func GetNonNegativeFixedDecs(precision uint32, decimals ...string) ([]math.Dec, error) {
	decs := make([]math.Dec, len(decimals))
	for i, decimal := range decimals {
		dec, err := math.NewNonNegativeFixedDecFromString(decimal, precision)
		if err != nil {
			return nil, err
		}
		decs[i] = dec
	}
	return decs, nil
}
