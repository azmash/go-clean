package shop

import "context"

type KYC interface {
	CheckKYC(ctx context.Context, shopID int64) (bool, error)
}

type ShopScore interface {
	GetShopScore(ctx context.Context, shopID int64) (int, error)
}

type Order interface {
	GetTotalFinishOrder(ctx context.Context, shopID int64) (int64, error)
	GetNetIncome(ctx context.Context, shopID int64) (int64, error)
}

type ShopClass interface {
	SetMembership(ctx context.Context, shopID int64, newStatus int) error
}
