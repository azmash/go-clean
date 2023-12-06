package grpc

import "context"

type ShopScore struct {
}

func NewGRPCShopScore() *ShopScore {
	return &ShopScore{}
}

func (s *ShopScore) GetShopScore(ctx context.Context, shopID int64) (int, error) {
	if shopID == 480824 {
		return 60, nil
	}

	return 0, nil
}
