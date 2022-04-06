package simulation

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	regentypes "github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
	basketsims "github.com/regen-network/regen-ledger/x/ecocredit/simulation/basket"
	marketplacesims "github.com/regen-network/regen-ledger/x/ecocredit/simulation/marketplace"
	"github.com/regen-network/regen-ledger/x/ecocredit/simulation/utils"
)

// Simulation operation weights constants
const (
	OpWeightMsgCreateClass         = "op_weight_msg_create_class"
	OpWeightMsgCreateBatch         = "op_weight_msg_create_batch"
	OpWeightMsgSend                = "op_weight_msg_send"
	OpWeightMsgRetire              = "op_weight_msg_retire"
	OpWeightMsgCancel              = "op_weight_msg_cancel"
	OpWeightMsgUpdateClassAdmin    = "op_weight_msg_update_class_admin"
	OpWeightMsgUpdateClassMetadata = "op_weight_msg_update_class_metadata"
	OpWeightMsgUpdateClassIssuers  = "op_weight_msg_update_class_issuers"
	OpWeightMsgCreateProject       = "op_weight_msg_create_project"
)

// ecocredit operations weights
const (
	WeightCreateClass   = 10
	WeightCreateProject = 20
	WeightCreateBatch   = 50
	WeightSend          = 100
	WeightRetire        = 80
	WeightCancel        = 30
	WeightUpdateClass   = 30
)

// ecocredit message types
var (
	TypeMsgCreateClass         = sdk.MsgTypeURL(&ecocredit.MsgCreateClass{})
	TypeMsgCreateProject       = sdk.MsgTypeURL(&ecocredit.MsgCreateProject{})
	TypeMsgCreateBatch         = sdk.MsgTypeURL(&ecocredit.MsgCreateBatch{})
	TypeMsgSend                = sdk.MsgTypeURL(&ecocredit.MsgSend{})
	TypeMsgRetire              = sdk.MsgTypeURL(&ecocredit.MsgRetire{})
	TypeMsgCancel              = sdk.MsgTypeURL(&ecocredit.MsgCancel{})
	TypeMsgUpdateClassAdmin    = sdk.MsgTypeURL(&ecocredit.MsgUpdateClassAdmin{})
	TypeMsgUpdateClassIssuers  = sdk.MsgTypeURL(&ecocredit.MsgUpdateClassIssuers{})
	TypeMsgUpdateClassMetadata = sdk.MsgTypeURL(&ecocredit.MsgUpdateClassMetadata{})
)

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec,
	ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient ecocredit.QueryClient, basketQryClient basket.QueryClient) simulation.WeightedOperations {

	var (
		weightMsgCreateClass         int
		weightMsgCreateBatch         int
		weightMsgSend                int
		weightMsgRetire              int
		weightMsgCancel              int
		weightMsgUpdateClassAdmin    int
		weightMsgUpdateClassIssuers  int
		weightMsgUpdateClassMetadata int
		weightMsgCreateProject       int
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgCreateClass, &weightMsgCreateClass, nil,
		func(_ *rand.Rand) {
			weightMsgCreateClass = WeightCreateClass
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgCreateBatch, &weightMsgCreateBatch, nil,
		func(_ *rand.Rand) {
			weightMsgCreateBatch = WeightCreateBatch
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgSend, &weightMsgSend, nil,
		func(_ *rand.Rand) {
			weightMsgSend = WeightSend
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgRetire, &weightMsgRetire, nil,
		func(_ *rand.Rand) {
			weightMsgRetire = WeightRetire
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgCancel, &weightMsgCancel, nil,
		func(_ *rand.Rand) {
			weightMsgCancel = WeightCancel
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgUpdateClassAdmin, &weightMsgUpdateClassAdmin, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateClassAdmin = WeightUpdateClass
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgUpdateClassIssuers, &weightMsgUpdateClassIssuers, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateClassIssuers = WeightUpdateClass
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgUpdateClassMetadata, &weightMsgUpdateClassMetadata, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateClassMetadata = WeightUpdateClass
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgCreateProject, &weightMsgCreateProject, nil,
		func(_ *rand.Rand) {
			weightMsgCreateProject = WeightCreateProject
		},
	)

	ops := simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgCreateClass,
			SimulateMsgCreateClass(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgCreateBatch,
			SimulateMsgCreateBatch(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgSend,
			SimulateMsgSend(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgRetire,
			SimulateMsgRetire(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgCancel,
			SimulateMsgCancel(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateClassAdmin,
			SimulateMsgUpdateClassAdmin(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateClassIssuers,
			SimulateMsgUpdateClassIssuers(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateClassMetadata,
			SimulateMsgUpdateClassMetadata(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgCreateProject,
			SimulateMsgCreateProject(ak, bk, qryClient),
		),
	}

	basketOps := basketsims.WeightedOperations(appParams, cdc, ak, bk, qryClient, basketQryClient)
	marketplaceOps := marketplacesims.WeightedOperations(appParams, cdc, ak, bk, qryClient, basketQryClient)

	ops = append(ops, basketOps...)
	return append(ops, marketplaceOps...)
}

// SimulateMsgCreateClass generates a MsgCreateClass with random values.
func SimulateMsgCreateClass(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient ecocredit.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		admin := accs[0]
		issuers := randomIssuers(r, accs)

		ctx := regentypes.Context{Context: sdkCtx}
		res, err := qryClient.Params(ctx, &ecocredit.QueryParamsRequest{})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateClass, err.Error()), nil, err
		}

		params := res.Params
		if params.AllowlistEnabled && !utils.Contains(params.AllowedClassCreators, admin.Address.String()) {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateClass, "not allowed to create credit class"), nil, nil // skip
		}

		spendable := bk.SpendableCoins(sdkCtx, admin.Address)
		if spendable.IsAllLTE(params.CreditClassFee) {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateClass, "not enough balance"), nil, nil
		}

		creditTypes := []string{"carbon", "biodiversity"}

		msg := &ecocredit.MsgCreateClass{
			Admin:          admin.Address.String(),
			Issuers:        issuers,
			Metadata:       []byte(simtypes.RandStringOfLength(r, 10)),
			CreditTypeName: creditTypes[r.Intn(len(creditTypes))],
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      admin,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateMsgCreateProject generates a MsgCreateProject with random values.
func SimulateMsgCreateProject(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient ecocredit.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		issuer := accs[0]

		classes, err := utils.GetAndShuffleClasses(sdkCtx, r, qryClient)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateProject, err.Error()), nil, err
		}

		if len(classes) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateProject, "no credit classes"), nil, nil
		}

		var classID string
		for _, class := range classes {
			if utils.Contains(class.Issuers, issuer.Address.String()) {
				classID = class.ClassId
				break
			}
		}

		if classID == "" {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateProject, "don't have permission to create project"), nil, nil
		}

		issuerAcc := ak.GetAccount(sdkCtx, issuer.Address)
		spendable := bk.SpendableCoins(sdkCtx, issuerAcc.GetAddress())

		msg := &ecocredit.MsgCreateProject{
			Issuer:          issuer.Address.String(),
			ClassId:         classID,
			Metadata:        []byte(simtypes.RandStringOfLength(r, 100)),
			ProjectLocation: "AB-CDE FG1 345",
		}
		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      issuer,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateMsgCreateBatch generates a MsgCreateBatch with random values.
func SimulateMsgCreateBatch(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient ecocredit.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		issuer := accs[0]

		ctx := regentypes.Context{Context: sdkCtx}
		res, err := qryClient.Projects(ctx, &ecocredit.QueryProjectsRequest{})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateBatch, err.Error()), nil, err
		}

		projects := res.Projects
		if len(projects) == 0 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateBatch, "no projects"), nil, nil
		}

		var projectID string
		for _, project := range projects {
			if project.Issuer == issuer.Address.String() {
				projectID = project.ProjectId
				break
			}
		}

		if projectID == "" {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCreateBatch, "don't have permission to create credit batch"), nil, nil
		}

		issuerAcc := ak.GetAccount(sdkCtx, issuer.Address)
		spendable := bk.SpendableCoins(sdkCtx, issuerAcc.GetAddress())

		now := sdkCtx.BlockTime()
		tenHours := now.Add(10 * time.Hour)
		msg := &ecocredit.MsgCreateBatch{
			Issuer:    issuer.Address.String(),
			ProjectId: projectID,
			Issuance:  generateBatchIssuance(r, accs),
			StartDate: &now,
			EndDate:   &tenHours,
			Metadata:  []byte(simtypes.RandStringOfLength(r, 10)),
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      issuer,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateMsgSend generates a MsgSend with random values.
func SimulateMsgSend(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient ecocredit.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		ctx := regentypes.Context{Context: sdkCtx}
		class, op, err := getRandomClass(sdkCtx, r, qryClient, TypeMsgSend)
		if class == nil {
			return op, nil, err
		}

		project, op, err := getRandomProjectFromClass(ctx, r, qryClient, TypeMsgSend, class.ClassId)
		if project == nil {
			return op, nil, err
		}

		batch, op, err := getRandomBatchFromProject(ctx, r, qryClient, TypeMsgSend, class.ClassId)
		if batch == nil {
			return op, nil, err
		}

		balres, err := qryClient.Balance(ctx, &ecocredit.QueryBalanceRequest{
			Account:    project.Issuer,
			BatchDenom: batch.BatchDenom,
		})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, err.Error()), nil, err
		}

		tradableBalance, err := math.NewNonNegativeDecFromString(balres.TradableAmount)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, err.Error()), nil, err
		}

		retiredBalance, err := math.NewNonNegativeDecFromString(balres.RetiredAmount)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, err.Error()), nil, err
		}

		if tradableBalance.IsZero() || retiredBalance.IsZero() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, "balance is zero"), nil, nil
		}

		recipient, _ := simtypes.RandomAcc(r, accs)
		if project.Issuer == recipient.Address.String() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, "sender & recipient are same"), nil, nil
		}

		addr, err := sdk.AccAddressFromBech32(project.Issuer)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, err.Error()), nil, err
		}

		acc, found := simtypes.FindAccount(accs, addr)
		if !found {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, "account not found"), nil, nil
		}

		issuer := ak.GetAccount(sdkCtx, acc.Address)
		spendable := bk.SpendableCoins(sdkCtx, issuer.GetAddress())

		var tradable int
		var retired int
		var retirementLocation string
		if !tradableBalance.IsZero() {
			i64, err := tradableBalance.Int64()
			if err != nil {
				return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, err.Error()), nil, nil
			}
			if i64 > 1 {
				tradable = simtypes.RandIntBetween(r, 1, int(i64))
				retired = simtypes.RandIntBetween(r, 0, tradable)
				if retired != 0 {
					retirementLocation = "AQ"
				}
			} else {
				tradable = int(i64)
			}
		}

		if retired+tradable > tradable {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, "insufficient credit balance"), nil, nil
		}

		msg := &ecocredit.MsgSend{
			Sender:    project.Issuer,
			Recipient: recipient.Address.String(),
			Credits: []*ecocredit.MsgSend_SendCredits{
				{
					BatchDenom:         batch.BatchDenom,
					TradableAmount:     fmt.Sprintf("%d", tradable),
					RetiredAmount:      fmt.Sprintf("%d", retired),
					RetirementLocation: retirementLocation,
				},
			},
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      acc,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateMsgRetire generates a MsgRetire with random values.
func SimulateMsgRetire(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient ecocredit.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		ctx := regentypes.Context{Context: sdkCtx}
		class, op, err := getRandomClass(sdkCtx, r, qryClient, TypeMsgRetire)
		if class == nil {
			return op, nil, err
		}

		project, op, err := getRandomProjectFromClass(ctx, r, qryClient, TypeMsgRetire, class.ClassId)
		if project == nil {
			return op, nil, err
		}

		batch, op, err := getRandomBatchFromProject(ctx, r, qryClient, TypeMsgRetire, project.ProjectId)
		if batch == nil {
			return op, nil, err
		}

		balanceRes, err := qryClient.Balance(ctx, &ecocredit.QueryBalanceRequest{
			Account:    project.Issuer,
			BatchDenom: batch.BatchDenom,
		})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, err.Error()), nil, err
		}

		tradableBalance, err := math.NewNonNegativeDecFromString(balanceRes.TradableAmount)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRetire, err.Error()), nil, err
		}

		if tradableBalance.IsZero() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRetire, "balance is zero"), nil, nil
		}

		randSub := math.NewDecFromInt64(int64(simtypes.RandIntBetween(r, 1, 10)))
		spendable, account, op, err := getAccountAndSpendableCoins(sdkCtx, bk, accs, project.Issuer, TypeMsgRetire)
		if spendable == nil {
			return op, nil, err
		}

		if !spendable.IsAllPositive() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgRetire, "insufficient funds"), nil, nil
		}

		if tradableBalance.Cmp(randSub) != 1 {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgSend, "insufficient funds"), nil, nil
		}

		msg := &ecocredit.MsgRetire{
			Holder: account.Address.String(),
			Credits: []*ecocredit.MsgRetire_RetireCredits{
				{
					BatchDenom: batch.BatchDenom,
					Amount:     randSub.String(),
				},
			},
			Location: "ST-UVW XY Z12",
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      *account,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateMsgCancel generates a MsgCancel with random values.
func SimulateMsgCancel(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient ecocredit.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {

		ctx := regentypes.Context{Context: sdkCtx}
		class, op, err := getRandomClass(sdkCtx, r, qryClient, TypeMsgCancel)
		if class == nil {
			return op, nil, err
		}

		project, op, err := getRandomProjectFromClass(ctx, r, qryClient, TypeMsgRetire, class.ClassId)
		if project == nil {
			return op, nil, err
		}

		batch, op, err := getRandomBatchFromProject(ctx, r, qryClient, TypeMsgCancel, project.ProjectId)
		if batch == nil {
			return op, nil, err
		}

		spendable, account, op, err := getAccountAndSpendableCoins(sdkCtx, bk, accs, project.Issuer, TypeMsgCancel)
		if spendable == nil {
			return op, nil, err
		}

		balanceRes, err := qryClient.Balance(ctx, &ecocredit.QueryBalanceRequest{
			Account:    project.Issuer,
			BatchDenom: batch.BatchDenom,
		})
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCancel, err.Error()), nil, err
		}

		tradableBalance, err := math.NewNonNegativeDecFromString(balanceRes.TradableAmount)
		if err != nil {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCancel, err.Error()), nil, err
		}

		if tradableBalance.IsZero() {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgCancel, "balance is zero"), nil, nil
		}

		msg := &ecocredit.MsgCancel{
			Holder: project.Issuer,
			Credits: []*ecocredit.MsgCancel_CancelCredits{
				{
					BatchDenom: batch.BatchDenom,
					Amount:     balanceRes.TradableAmount,
				},
			},
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      *account,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateMsgUpdateClassAdmin generates a MsgUpdateClassAdmin with random values
func SimulateMsgUpdateClassAdmin(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient ecocredit.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		class, op, err := getRandomClass(sdkCtx, r, qryClient, TypeMsgUpdateClassAdmin)
		if class == nil {
			return op, nil, err
		}

		spendable, account, op, err := getAccountAndSpendableCoins(sdkCtx, bk, accs, class.Admin, TypeMsgUpdateClassAdmin)
		if spendable == nil {
			return op, nil, err
		}

		newAdmin, _ := simtypes.RandomAcc(r, accs)
		if newAdmin.Address.String() == class.Admin {
			return simtypes.NoOpMsg(ecocredit.ModuleName, TypeMsgUpdateClassAdmin, "same account"), nil, nil // skip
		}

		msg := &ecocredit.MsgUpdateClassAdmin{
			Admin:    class.Admin,
			ClassId:  class.ClassId,
			NewAdmin: newAdmin.Address.String(),
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      *account,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateMsgUpdateClassMetadata generates a MsgUpdateClassMetadata with random metadata
func SimulateMsgUpdateClassMetadata(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient ecocredit.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		class, op, err := getRandomClass(sdkCtx, r, qryClient, TypeMsgUpdateClassMetadata)
		if class == nil {
			return op, nil, err
		}

		spendable, account, op, err := getAccountAndSpendableCoins(sdkCtx, bk, accs, class.Admin, TypeMsgUpdateClassMetadata)
		if spendable == nil {
			return op, nil, err
		}

		msg := &ecocredit.MsgUpdateClassMetadata{
			Admin:    class.Admin,
			ClassId:  class.ClassId,
			Metadata: []byte(simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 10, 256))),
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      *account,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(txCtx)
	}
}

// SimulateMsgUpdateClassIssuers generates a MsgUpdateClassMetaData with random values
func SimulateMsgUpdateClassIssuers(ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	qryClient ecocredit.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		class, op, err := getRandomClass(sdkCtx, r, qryClient, TypeMsgUpdateClassIssuers)
		if class == nil {
			return op, nil, err
		}

		spendable, account, op, err := getAccountAndSpendableCoins(sdkCtx, bk, accs, class.Admin, TypeMsgUpdateClassIssuers)
		if spendable == nil {
			return op, nil, err
		}

		issuers := randomIssuers(r, accs)
		msg := &ecocredit.MsgUpdateClassIssuers{
			Admin:   class.Admin,
			ClassId: class.ClassId,
			Issuers: issuers,
		}

		txCtx := simulation.OperationInput{
			R:               r,
			App:             app,
			TxGen:           simappparams.MakeTestEncodingConfig().TxConfig,
			Cdc:             nil,
			Msg:             msg,
			MsgType:         msg.Type(),
			Context:         sdkCtx,
			SimAccount:      *account,
			AccountKeeper:   ak,
			Bankkeeper:      bk,
			ModuleName:      ecocredit.ModuleName,
			CoinsSpentInMsg: spendable,
		}

		return utils.GenAndDeliverTxWithRandFees(txCtx)
	}
}

func getAccountAndSpendableCoins(ctx sdk.Context, bk ecocredit.BankKeeper,
	accs []simtypes.Account, addr, msgType string) (sdk.Coins, *simtypes.Account, simtypes.OperationMsg, error) {
	accAddr, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return nil, nil, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, err.Error()), err
	}

	account, found := simtypes.FindAccount(accs, accAddr)
	if !found {
		return nil, &account, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, "account not found"), nil
	}

	spendable := bk.SpendableCoins(ctx, accAddr)
	return spendable, &account, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, ""), nil

}

func getRandomClass(ctx sdk.Context, r *rand.Rand, qryClient ecocredit.QueryClient, msgType string) (*ecocredit.ClassInfo, simtypes.OperationMsg, error) {
	classes, err := utils.GetAndShuffleClasses(ctx, r, qryClient)
	if err != nil {
		return nil, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, err.Error()), err
	}

	if len(classes) == 0 {
		return nil, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, "no credit class found"), nil
	}

	return classes[0], simtypes.NoOpMsg(ecocredit.ModuleName, msgType, ""), nil
}

func getRandomProjectFromClass(ctx regentypes.Context, r *rand.Rand, qryClient ecocredit.QueryClient, msgType, classID string) (*ecocredit.ProjectInfo, simtypes.OperationMsg, error) {
	res, err := qryClient.Projects(ctx, &ecocredit.QueryProjectsRequest{
		ClassId: classID,
	})
	if err != nil {
		return nil, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, err.Error()), err
	}

	projects := res.Projects
	if len(projects) == 0 {
		return nil, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, "no project found"), nil
	}

	return projects[r.Intn(len(projects))], simtypes.NoOpMsg(ecocredit.ModuleName, msgType, ""), nil
}

func getRandomBatchFromProject(ctx regentypes.Context, r *rand.Rand, qryClient ecocredit.QueryClient, msgType, projectID string) (*ecocredit.BatchInfo, simtypes.OperationMsg, error) {
	res, err := qryClient.Batches(ctx, &ecocredit.QueryBatchesRequest{
		ProjectId: projectID,
	})
	if err != nil {
		return nil, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, err.Error()), err
	}

	batches := res.Batches
	if len(batches) == 0 {
		return nil, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, "no batch found"), nil
	}

	return batches[r.Intn(len(batches))], simtypes.NoOpMsg(ecocredit.ModuleName, msgType, ""), nil
}

func randomIssuers(r *rand.Rand, accounts []simtypes.Account) []string {
	n := simtypes.RandIntBetween(r, 3, 10)
	issuers := make([]string, n)
	for i := 0; i < n; i++ {
		acc, _ := simtypes.RandomAcc(r, accounts)
		issuers[i] = acc.Address.String()
	}

	return issuers
}

func generateBatchIssuance(r *rand.Rand, accs []simtypes.Account) []*ecocredit.MsgCreateBatch_BatchIssuance {
	numIssuances := simtypes.RandIntBetween(r, 3, 10)
	res := make([]*ecocredit.MsgCreateBatch_BatchIssuance, numIssuances)

	for i := 0; i < numIssuances; i++ {
		recipient := accs[i]
		retiredAmount := simtypes.RandIntBetween(r, 0, 100)
		var retirementLocation string
		if retiredAmount > 0 {
			retirementLocation = "AD"
		}
		res[i] = &ecocredit.MsgCreateBatch_BatchIssuance{
			Recipient:          recipient.Address.String(),
			TradableAmount:     fmt.Sprintf("%d", simtypes.RandIntBetween(r, 10, 1000)),
			RetiredAmount:      fmt.Sprintf("%d", retiredAmount),
			RetirementLocation: retirementLocation,
		}
	}

	return res
}
