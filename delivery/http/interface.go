package http

import "context"

type Usecases interface {
	SetPM(ctx context.Context, shopID int64) (newStatus int, err error)
}
