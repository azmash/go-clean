package shop

type Repo struct {
	KYC
	ShopScore
	Order
	ShopClass
}

func NewUsecase(repoKYC KYC, repoShopScore ShopScore, repoOrder Order, repoShopClass ShopClass) *Repo {
	return &Repo{
		KYC:       repoKYC,
		ShopScore: repoShopScore,
		Order:     repoOrder,
		ShopClass: repoShopClass,
	}
}
