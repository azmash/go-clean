package grpc

import "context"

type Kyc struct {
}

func NewGRPCKYC() *Kyc {
	return &Kyc{}
}

func (k *Kyc) CheckKYC(ctx context.Context, shopID int64) (bool, error) {
	if shopID == 480824 {
		return true, nil
	}

	return false, nil
}
