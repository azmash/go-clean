package grpc

type GRPC struct {
	*Kyc
	*ShopScore
}

func NewGRPC() *GRPC {
	return &GRPC{}
}
