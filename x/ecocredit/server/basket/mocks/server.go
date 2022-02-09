// Code generated by MockGen. DO NOT EDIT.
// Source: x/ecocredit/server/basket/server.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	gomock "github.com/golang/mock/gomock"
	ecocredit "github.com/regen-network/regen-ledger/x/ecocredit"
)

// MockEcocreditKeeper is a mock of EcocreditKeeper interface.
type MockEcocreditKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockEcocreditKeeperMockRecorder
}

// MockEcocreditKeeperMockRecorder is the mock recorder for MockEcocreditKeeper.
type MockEcocreditKeeperMockRecorder struct {
	mock *MockEcocreditKeeper
}

// NewMockEcocreditKeeper creates a new mock instance.
func NewMockEcocreditKeeper(ctrl *gomock.Controller) *MockEcocreditKeeper {
	mock := &MockEcocreditKeeper{ctrl: ctrl}
	mock.recorder = &MockEcocreditKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEcocreditKeeper) EXPECT() *MockEcocreditKeeperMockRecorder {
	return m.recorder
}

// Balance mocks base method.
func (m *MockEcocreditKeeper) Balance(arg0 context.Context, arg1 *ecocredit.QueryBalanceRequest) (*ecocredit.QueryBalanceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Balance", arg0, arg1)
	ret0, _ := ret[0].(*ecocredit.QueryBalanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Balance indicates an expected call of Balance.
func (mr *MockEcocreditKeeperMockRecorder) Balance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Balance", reflect.TypeOf((*MockEcocreditKeeper)(nil).Balance), arg0, arg1)
}

// BatchInfo mocks base method.
func (m *MockEcocreditKeeper) BatchInfo(arg0 context.Context, arg1 *ecocredit.QueryBatchInfoRequest) (*ecocredit.QueryBatchInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BatchInfo", arg0, arg1)
	ret0, _ := ret[0].(*ecocredit.QueryBatchInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BatchInfo indicates an expected call of BatchInfo.
func (mr *MockEcocreditKeeperMockRecorder) BatchInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BatchInfo", reflect.TypeOf((*MockEcocreditKeeper)(nil).BatchInfo), arg0, arg1)
}

// Batches mocks base method.
func (m *MockEcocreditKeeper) Batches(arg0 context.Context, arg1 *ecocredit.QueryBatchesRequest) (*ecocredit.QueryBatchesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Batches", arg0, arg1)
	ret0, _ := ret[0].(*ecocredit.QueryBatchesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Batches indicates an expected call of Batches.
func (mr *MockEcocreditKeeperMockRecorder) Batches(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Batches", reflect.TypeOf((*MockEcocreditKeeper)(nil).Batches), arg0, arg1)
}

// ClassInfo mocks base method.
func (m *MockEcocreditKeeper) ClassInfo(arg0 context.Context, arg1 *ecocredit.QueryClassInfoRequest) (*ecocredit.QueryClassInfoResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClassInfo", arg0, arg1)
	ret0, _ := ret[0].(*ecocredit.QueryClassInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ClassInfo indicates an expected call of ClassInfo.
func (mr *MockEcocreditKeeperMockRecorder) ClassInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClassInfo", reflect.TypeOf((*MockEcocreditKeeper)(nil).ClassInfo), arg0, arg1)
}

// Classes mocks base method.
func (m *MockEcocreditKeeper) Classes(arg0 context.Context, arg1 *ecocredit.QueryClassesRequest) (*ecocredit.QueryClassesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Classes", arg0, arg1)
	ret0, _ := ret[0].(*ecocredit.QueryClassesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Classes indicates an expected call of Classes.
func (mr *MockEcocreditKeeperMockRecorder) Classes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Classes", reflect.TypeOf((*MockEcocreditKeeper)(nil).Classes), arg0, arg1)
}

// CreditTypes mocks base method.
func (m *MockEcocreditKeeper) CreditTypes(arg0 context.Context, arg1 *ecocredit.QueryCreditTypesRequest) (*ecocredit.QueryCreditTypesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreditTypes", arg0, arg1)
	ret0, _ := ret[0].(*ecocredit.QueryCreditTypesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreditTypes indicates an expected call of CreditTypes.
func (mr *MockEcocreditKeeperMockRecorder) CreditTypes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreditTypes", reflect.TypeOf((*MockEcocreditKeeper)(nil).CreditTypes), arg0, arg1)
}

// GetCreateBasketFee mocks base method.
func (m *MockEcocreditKeeper) GetCreateBasketFee(ctx context.Context) types.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCreateBasketFee", ctx)
	ret0, _ := ret[0].(types.Coins)
	return ret0
}

// GetCreateBasketFee indicates an expected call of GetCreateBasketFee.
func (mr *MockEcocreditKeeperMockRecorder) GetCreateBasketFee(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCreateBasketFee", reflect.TypeOf((*MockEcocreditKeeper)(nil).GetCreateBasketFee), ctx)
}

// Params mocks base method.
func (m *MockEcocreditKeeper) Params(arg0 context.Context, arg1 *ecocredit.QueryParamsRequest) (*ecocredit.QueryParamsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Params", arg0, arg1)
	ret0, _ := ret[0].(*ecocredit.QueryParamsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Params indicates an expected call of Params.
func (mr *MockEcocreditKeeperMockRecorder) Params(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Params", reflect.TypeOf((*MockEcocreditKeeper)(nil).Params), arg0, arg1)
}

// Supply mocks base method.
func (m *MockEcocreditKeeper) Supply(arg0 context.Context, arg1 *ecocredit.QuerySupplyRequest) (*ecocredit.QuerySupplyResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Supply", arg0, arg1)
	ret0, _ := ret[0].(*ecocredit.QuerySupplyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Supply indicates an expected call of Supply.
func (mr *MockEcocreditKeeperMockRecorder) Supply(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Supply", reflect.TypeOf((*MockEcocreditKeeper)(nil).Supply), arg0, arg1)
}

// MockBankKeeper is a mock of BankKeeper interface.
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
}

// MockBankKeeperMockRecorder is the mock recorder for MockBankKeeper.
type MockBankKeeperMockRecorder struct {
	mock *MockBankKeeper
}

// NewMockBankKeeper creates a new mock instance.
func NewMockBankKeeper(ctrl *gomock.Controller) *MockBankKeeper {
	mock := &MockBankKeeper{ctrl: ctrl}
	mock.recorder = &MockBankKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankKeeper) EXPECT() *MockBankKeeperMockRecorder {
	return m.recorder
}

// BurnCoins mocks base method.
func (m *MockBankKeeper) BurnCoins(ctx types.Context, moduleName string, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BurnCoins", ctx, moduleName, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// BurnCoins indicates an expected call of BurnCoins.
func (mr *MockBankKeeperMockRecorder) BurnCoins(ctx, moduleName, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BurnCoins", reflect.TypeOf((*MockBankKeeper)(nil).BurnCoins), ctx, moduleName, amt)
}

// MintCoins mocks base method.
func (m *MockBankKeeper) MintCoins(ctx types.Context, moduleName string, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MintCoins", ctx, moduleName, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// MintCoins indicates an expected call of MintCoins.
func (mr *MockBankKeeperMockRecorder) MintCoins(ctx, moduleName, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MintCoins", reflect.TypeOf((*MockBankKeeper)(nil).MintCoins), ctx, moduleName, amt)
}

// SendCoinsFromAccountToModule mocks base method.
func (m *MockBankKeeper) SendCoinsFromAccountToModule(ctx types.Context, senderAddr types.AccAddress, recipientModule string, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromAccountToModule", ctx, senderAddr, recipientModule, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromAccountToModule indicates an expected call of SendCoinsFromAccountToModule.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromAccountToModule(ctx, senderAddr, recipientModule, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromAccountToModule", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromAccountToModule), ctx, senderAddr, recipientModule, amt)
}

// SendCoinsFromModuleToAccount mocks base method.
func (m *MockBankKeeper) SendCoinsFromModuleToAccount(ctx types.Context, senderModule string, recipientAddr types.AccAddress, amt types.Coins) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendCoinsFromModuleToAccount", ctx, senderModule, recipientAddr, amt)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendCoinsFromModuleToAccount indicates an expected call of SendCoinsFromModuleToAccount.
func (mr *MockBankKeeperMockRecorder) SendCoinsFromModuleToAccount(ctx, senderModule, recipientAddr, amt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendCoinsFromModuleToAccount", reflect.TypeOf((*MockBankKeeper)(nil).SendCoinsFromModuleToAccount), ctx, senderModule, recipientAddr, amt)
}