package core

import (
	"context"
	"github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/mock/gomock"
	ecocreditv1beta1 "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1beta1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/v1beta1"
	"gotest.tools/v3/assert"
	"testing"
	"time"
)

func TestValidBatch(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	before(t, s.ctx, s.stateStore, s.addr)

	any := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(any, any).Do(func(any, p *ecocredit.Params)  {
		p.AllowlistEnabled = false
		p.CreditClassFee = types.NewCoins(types.NewInt64Coin("foo", 20))
		p.CreditTypes = []*ecocredit.CreditType{{Name: "carbon", Abbreviation: "C", Unit: "tonne", Precision: 6}}
	}).AnyTimes()

	start, end := time.Now(), time.Now()
	res, err := s.k.CreateBatch(s.ctx, &v1beta1.MsgCreateBatch{
		Issuer:    s.addr.String(),
		ProjectId: "PRO",
		Issuance:  []*v1beta1.MsgCreateBatch_BatchIssuance{
			{
				Recipient: s.addr.String(),
				TradableAmount: "10",
				RetiredAmount: "5.3",
			},
		},
		Metadata:  nil,
		StartDate: &start,
		EndDate:   &end,
	})

	batch, err := s.stateStore.BatchInfoStore().Get(s.ctx, 1)
	assert.NilError(t, err)
	assert.Equal(t, res.BatchDenom, batch.BatchDenom)

	assert.NilError(t, err, "unexpected error: %w", err)
	sup, err := s.stateStore.BatchSupplyStore().Get(s.ctx, 1)
	assert.NilError(t, err)
	assert.Equal(t, "10", sup.TradableAmount, "got %s", sup.TradableAmount)
	assert.Equal(t, "5.3", sup.RetiredAmount, "got %s", sup.RetiredAmount)
	assert.Equal(t, "0", sup.CancelledAmount, "got %s", sup.CancelledAmount)

	bal, err := s.stateStore.BatchBalanceStore().Get(s.ctx, s.addr, 1)
	assert.NilError(t, err)
	assert.Equal(t, "10", bal.Tradable)
	assert.Equal(t, "5.3", bal.Retired)
}

func TestBadPrecision(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	before(t, s.ctx, s.stateStore, s.addr)

	any := gomock.Any()
	s.paramsKeeper.EXPECT().GetParamSet(any, any).Do(func(any, p *ecocredit.Params)  {
		p.AllowlistEnabled = false
		p.CreditClassFee = types.NewCoins(types.NewInt64Coin("foo", 20))
		p.CreditTypes = []*ecocredit.CreditType{{Name: "carbon", Abbreviation: "C", Unit: "tonne", Precision: 6}}
	}).AnyTimes()

	start, end := time.Now(), time.Now()
	_, err := s.k.CreateBatch(s.ctx, &v1beta1.MsgCreateBatch{
		Issuer:    s.addr.String(),
		ProjectId: "PRO",
		Issuance:  []*v1beta1.MsgCreateBatch_BatchIssuance{
			{
				Recipient: s.addr.String(),
				TradableAmount: "10.1234567891111",
			},
		},
		Metadata:  nil,
		StartDate: &start,
		EndDate:   &end,
	})
	assert.ErrorContains(t, err, "exceeds maximum decimal places")
}

func TestUnauthorizedIssuer(t *testing.T) {
	t.Parallel()
	s := setupBase(t)
	before(t, s.ctx, s.stateStore, s.addr)
	_, err := s.k.CreateBatch(s.ctx, &v1beta1.MsgCreateBatch{
		ProjectId: "PRO",
		Issuer: "FooBarBaz",
	})
	assert.ErrorContains(t, err, "is not an issuer for the class")
}

func TestProjectNotFound(t *testing.T) {
	t.Parallel()
	s := setupBase(t)

	_, err := s.k.CreateBatch(s.ctx, &v1beta1.MsgCreateBatch{
		ProjectId: "none",
	})
	assert.ErrorContains(t, err, "not found")
}

func before(t *testing.T, ctx context.Context, ss ecocreditv1beta1.StateStore, addr types.AccAddress) (classRowId, projectRowId uint64){
	cid, err := ss.ClassInfoStore().InsertReturningID(ctx, &ecocreditv1beta1.ClassInfo{
		Name:       "C01",
		Admin:      addr,
		Metadata:   nil,
		CreditType: "C",
	})
	assert.NilError(t, err)
	err = ss.ClassIssuerStore().Insert(ctx, &ecocreditv1beta1.ClassIssuer{
		ClassId: cid,
		Issuer:  addr,
	})
	assert.NilError(t, err)
	pid, err := ss.ProjectInfoStore().InsertReturningID(ctx, &ecocreditv1beta1.ProjectInfo{
		Name:            "PRO",
		ClassId:         1,
		ProjectLocation: "",
		Metadata:        nil,
	})
	assert.NilError(t, err)

	classRowId, projectRowId = cid, pid
	return
}
