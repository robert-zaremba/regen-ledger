package core

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/types/query"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestQuery_Batches(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	// make a project and two batches
	assert.NilError(t, s.stateStore.ProjectTable().Insert(s.ctx, &api.Project{
		Id:                  "P01",
		ClassKey:            1,
		ProjectJurisdiction: "US-CA",
		Metadata:            "",
	}))
	assert.NilError(t, s.stateStore.BatchTable().Insert(s.ctx, &api.Batch{
		ProjectKey: 1,
		Denom:      "C01-20200101-20220101-001",
		Metadata:   "",
		StartDate:  nil,
		EndDate:    nil,
	}))
	assert.NilError(t, s.stateStore.BatchTable().Insert(s.ctx, &api.Batch{
		ProjectKey: 1,
		Denom:      "C01-20200101-20220101-002",
		Metadata:   "",
		StartDate:  nil,
		EndDate:    nil,
	}))

	// invalid query
	_, err := s.k.Batches(s.ctx, &core.QueryBatchesRequest{ProjectId: "F01"})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())

	//  base query
	res, err := s.k.Batches(s.ctx, &core.QueryBatchesRequest{ProjectId: "P01"})
	assert.NilError(t, err)
	assert.Equal(t, 2, len(res.Batches))
	assert.Equal(t, "C01-20200101-20220101-001", res.Batches[0].Denom)

	// paginated query
	res, err = s.k.Batches(s.ctx, &core.QueryBatchesRequest{
		ProjectId:  "P01",
		Pagination: &query.PageRequest{Limit: 1, CountTotal: true},
	})
	assert.NilError(t, err)
	assert.Equal(t, 1, len(res.Batches))
	assert.Equal(t, uint64(2), res.Pagination.Total)
}
