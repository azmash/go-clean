package http

import "github.com/azmash/go-clean/config"

type HTTP struct {
	*ShopMembership
	*Order
	*Kyc
	*ShopScore
}

func NewHTTP(url config.HostConfig) *HTTP {
	return &HTTP{
		ShopMembership: NewHTTPShopMembership(url.Host),
		Order:          NewHTTPOrder(url.Host),
		Kyc:            NewHTTPKYC(url.Host),
		ShopScore:      NewHTTPShopScore(url.Host),
	}
}
