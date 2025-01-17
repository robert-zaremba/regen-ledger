package core

import (
	"context"

	"github.com/regen-network/regen-ledger/types/ormutil"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// ClassInfo queries for information on a credit class.
func (k Keeper) ClassInfo(ctx context.Context, request *core.QueryClassInfoRequest) (*core.QueryClassInfoResponse, error) {
	classInfo, err := k.stateStore.ClassTable().GetById(ctx, request.ClassId)
	if err != nil {
		return nil, err
	}

	var ci core.Class
	if err = ormutil.PulsarToGogoSlow(classInfo, &ci); err != nil {
		return nil, err
	}
	return &core.QueryClassInfoResponse{Class: &ci}, nil
}
