package core

import (
	"testing"

	"gotest.tools/v3/assert"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func TestQuery_Balance(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	noBalanceAddr := genAddrs(1)[0]
	batchDenom := "C01-20200101-20220101-001"
	tradable := "10.54321"
	retired := "50.3214"

	// make a batch and give s.addr some balance
	assert.NilError(t, s.stateStore.BatchInfoTable().Insert(s.ctx, &api.BatchInfo{
		BatchDenom: batchDenom,
	}))
	assert.NilError(t, s.stateStore.BatchBalanceTable().Insert(s.ctx, &api.BatchBalance{
		Address:  s.addr,
		BatchId:  1,
		Tradable: tradable,
		Retired:  retired,
	}))

	// valid query
	res, err := s.k.Balance(s.ctx, &core.QueryBalanceRequest{
		Account:    s.addr.String(),
		BatchDenom: batchDenom,
	})
	assert.NilError(t, err)
	assert.Equal(t, s.addr.String(), res.Balance.Address)
	assert.Equal(t, batchDenom, res.Balance.BatchDenom)
	assert.Equal(t, tradable, res.Balance.Tradable)
	assert.Equal(t, retired, res.Balance.Retired)

	// random addr should just give 0
	res, err = s.k.Balance(s.ctx, &core.QueryBalanceRequest{
		Account:    noBalanceAddr.String(),
		BatchDenom: batchDenom,
	})
	assert.NilError(t, err)
	assert.Equal(t, noBalanceAddr.String(), res.Balance.Address)
	assert.Equal(t, batchDenom, res.Balance.BatchDenom)
	assert.Equal(t, "0", res.Balance.Tradable)
	assert.Equal(t, "0", res.Balance.Retired)

	// query with invalid batch should return not found
	_, err = s.k.Balance(s.ctx, &core.QueryBalanceRequest{
		Account:    s.addr.String(),
		BatchDenom: "A00-00000000-00000000-001",
	})
	assert.ErrorContains(t, err, ormerrors.NotFound.Error())
}
