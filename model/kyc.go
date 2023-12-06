package model

type KYCDetail struct {
	ShopID    int64 `json:"shop_id"`
	KycStatus int   `json:"kyc_status"`
}

type KYCResult struct {
	Code   int       `json:"code"`
	Status string    `json:"status"`
	Data   KYCDetail `json:"data"`
}

const KYCVerified int = 1
