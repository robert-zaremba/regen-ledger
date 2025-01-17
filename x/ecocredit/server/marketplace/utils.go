package marketplace

import (
	"context"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/marketplace/v1"
	ecoApi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

// isDenomAllowed checks if the denom is allowed to be used in orders.
func isDenomAllowed(ctx sdk.Context, denom string, pk ecocredit.ParamKeeper) bool {
	var params core.Params
	pk.GetParamSet(ctx, &params)
	for _, askDenom := range params.AllowedAskDenoms {
		if askDenom.Denom == denom {
			return true
		}
	}
	return false
}

type orderOptions struct {
	autoRetire         bool
	canPartialFill     bool
	batchDenom         string
	retirementLocation string
}

// fillOrder moves credits and coins according to the order. It will:
// - update a sell order, removing it if quantity becomes 0 as a result of this purchase.
// - remove the purchaseQty from the seller's escrowed balance.
// - add credits to the buyer's tradable/retired address (based on the DisableAutoRetire field).
// - update the supply accordingly.
// - send the coins specified in the bid to the seller.
func (k Keeper) fillOrder(ctx context.Context, sellOrder *api.SellOrder, buyerAcc sdk.AccAddress, purchaseQty math.Dec,
	cost sdk.Coin, opts orderOptions) error {
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sellOrderQty, err := math.NewDecFromString(sellOrder.Quantity)
	if err != nil {
		return err
	}

	switch sellOrderQty.Cmp(purchaseQty) {
	case math.LessThan:
		if !opts.canPartialFill {
			return ecocredit.ErrInsufficientCredits.Wrapf("cannot purchase %v credits from a sell order that has %s credits", purchaseQty, sellOrder.Quantity)
		} else {
			// if we can partially fill, we just delete the sellOrder and take whatever
			// credits are left from that order.
			if err := k.stateStore.SellOrderTable().Delete(ctx, sellOrder); err != nil {
				return err
			}
			purchaseQty = sellOrderQty
		}
	case math.EqualTo:
		if err := k.stateStore.SellOrderTable().Delete(ctx, sellOrder); err != nil {
			return err
		}
	case math.GreaterThan:
		newSellOrderQty, err := sellOrderQty.Sub(purchaseQty)
		if err != nil {
			return err
		}
		sellOrder.Quantity = newSellOrderQty.String()
		if err = k.stateStore.SellOrderTable().Update(ctx, sellOrder); err != nil {
			return err
		}
	}

	// remove the credits from the seller's escrowed balance
	sellerBal, err := k.coreStore.BatchBalanceTable().Get(ctx, sellOrder.Seller, sellOrder.BatchId)
	if err != nil {
		return err
	}
	escrowBal, err := math.NewDecFromString(sellerBal.Escrowed)
	if err != nil {
		return err
	}
	escrowBal, err = math.SafeSubBalance(escrowBal, purchaseQty)
	if err != nil {
		return err
	}
	sellerBal.Escrowed = escrowBal.String()
	if err = k.coreStore.BatchBalanceTable().Update(ctx, sellerBal); err != nil {
		return err
	}

	// update the buyers balance and the batch supply
	supply, err := k.coreStore.BatchSupplyTable().Get(ctx, sellOrder.BatchId)
	if err != nil {
		return err
	}
	buyerBal, err := k.coreStore.BatchBalanceTable().Get(ctx, buyerAcc, sellOrder.BatchId)
	if err != nil {
		if ormerrors.IsNotFound(err) {
			buyerBal = &ecoApi.BatchBalance{
				BatchKey: sellOrder.BatchId,
				Address:  buyerAcc,
				Tradable: "0",
				Retired:  "0",
				Escrowed: "0",
			}
		} else {
			return err
		}
	}
	if !opts.autoRetire { // if auto retire is disabled, we move the credits into the buyer's/supply's tradable balance.
		tradableBalance, err := math.NewDecFromString(buyerBal.Tradable)
		if err != nil {
			return err
		}
		tradableBalance, err = math.SafeAddBalance(tradableBalance, purchaseQty)
		if err != nil {
			return err
		}
		buyerBal.Tradable = tradableBalance.String()

		supplyTradable, err := math.NewDecFromString(supply.TradableAmount)
		if err != nil {
			return err
		}
		supplyTradable, err = math.SafeAddBalance(supplyTradable, purchaseQty)
		if err != nil {
			return err
		}
		supply.TradableAmount = supplyTradable.String()
		if err = sdkCtx.EventManager().EmitTypedEvent(&core.EventReceive{
			Sender:         sdk.AccAddress(sellOrder.Seller).String(),
			Recipient:      buyerAcc.String(),
			BatchDenom:     opts.batchDenom,
			TradableAmount: purchaseQty.String(),
			RetiredAmount:  "0",
		}); err != nil {
			return err
		}
	} else {
		retiredBalance, err := math.NewDecFromString(buyerBal.Retired)
		if err != nil {
			return err
		}
		retiredBalance, err = math.SafeAddBalance(retiredBalance, purchaseQty)
		if err != nil {
			return err
		}
		buyerBal.Retired = retiredBalance.String()

		supplyRetired, err := math.NewDecFromString(supply.RetiredAmount)
		if err != nil {
			return err
		}
		supplyRetired, err = math.SafeAddBalance(supplyRetired, purchaseQty)
		if err != nil {
			return err
		}
		supply.RetiredAmount = supplyRetired.String()
		if err = sdkCtx.EventManager().EmitTypedEvent(&core.EventRetire{
			Retirer:      buyerAcc.String(),
			BatchDenom:   opts.batchDenom,
			Amount:       purchaseQty.String(),
			Jurisdiction: opts.retirementLocation,
		}); err != nil {
			return err
		}
	}
	if err = k.coreStore.BatchSupplyTable().Update(ctx, supply); err != nil {
		return err
	}
	if err = k.coreStore.BatchBalanceTable().Save(ctx, buyerBal); err != nil {
		return err
	}

	return k.bankKeeper.SendCoins(sdkCtx, buyerAcc, sellOrder.Seller, sdk.NewCoins(cost))
}

// getTotalCost calculates the cost of the order by multiplying the price per credit, and the amount of credits
// desired in the order.
func getTotalCost(pricePerCredit sdk.Int, amtCredits math.Dec) (sdk.Int, error) {
	unitPrice, err := math.NewPositiveFixedDecFromString(pricePerCredit.String(), amtCredits.NumDecimalPlaces())
	if err != nil {
		return sdk.Int{}, err
	}
	cost, err := amtCredits.Mul(unitPrice)
	if err != nil {
		return sdk.Int{}, err
	}
	return cost.SdkIntTrim(), nil
}
