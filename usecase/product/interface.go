package product

import "context"

type GetProduct interface {
	GetProduct(ctx context.Context) (int64, error)
}
