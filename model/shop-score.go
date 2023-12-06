package model

type ShopScoreDetail struct {
	ShopID    int64 `json:"shop_id"`
	ShopScore int   `json:"shop_score"`
}

type ShopScoreResult struct {
	Code   int             `json:"code"`
	Status string          `json:"status"`
	Data   ShopScoreDetail `json:"data"`
}

const MinimumScorePM int = 60
const MinimumScorePMPro int = 70
