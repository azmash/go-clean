package model

type NetIncomeDetail struct {
	ShopID    int64 `json:"shop_id"`
	NetIncome int64 `json:"net_income"`
}

type FinishOrderDetail struct {
	ShopID      int64 `json:"shop_id"`
	FinishOrder int64 `json:"finish_order"`
}

type NetIncomeResult struct {
	Code   int             `json:"code"`
	Status string          `json:"status"`
	Data   NetIncomeDetail `json:"data"`
}

type FinishOrderResult struct {
	Code   int               `json:"code"`
	Status string            `json:"status"`
	Data   FinishOrderDetail `json:"data"`
}

const MinimumFinishOrder int = 3
const MinimumNetIncome int = 350000
