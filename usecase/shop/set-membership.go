package shop

import (
	"context"
	"errors"

	"github.com/azmash/go-clean/model"
)

func (r *Repo) SetPM(ctx context.Context, shopID int64) (newStatus int, err error) {
	isKYC, err := r.KYC.CheckKYC(ctx, shopID)
	if err != nil {
		return 0, err
	}

	if !isKYC {
		return 0, errors.New("shop is not yet kyc")
	}

	shopScore, err := r.ShopScore.GetShopScore(ctx, shopID)
	if err != nil {
		return 0, err
	}

	if shopScore >= model.MinimumScorePM {
		// Shop eligible to be PM/ Pro
		netIncome, err := r.Order.GetNetIncome(ctx, shopID)
		if err != nil {
			return 0, err
		}

		finishOrder, err := r.Order.GetTotalFinishOrder(ctx, shopID)
		if err != nil {
			return 0, err
		}

		if shopScore >= model.MinimumScorePMPro && finishOrder >= int64(model.MinimumFinishOrder) && netIncome >= int64(model.MinimumNetIncome) {
			newStatus = model.PowerMerchantPRO
		} else {
			newStatus = model.PowerMerchant
		}

		err = r.SetMembership(ctx, shopID, newStatus)
		if err != nil {
			return 0, err
		}

	}
	return newStatus, nil
}
