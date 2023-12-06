package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/azmash/go-clean/model"
)

type Order struct {
	host string
}

func NewHTTPOrder(host string) *Order {
	return &Order{
		host: host,
	}
}

func (t *Order) GetTotalFinishOrder(ctx context.Context, shopID int64) (int64, error) {
	var res model.FinishOrderResult
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/order/finish?shop_id=%d", t.host, shopID), nil)

	if err != nil {
		return 0, err
	}

	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	defer resp.Body.Close()

	err = json.Unmarshal(buf, &res)
	if err != nil {
		return 0, err
	}

	return res.Data.FinishOrder, nil
}

func (t *Order) GetNetIncome(ctx context.Context, shopID int64) (int64, error) {
	var res model.NetIncomeResult
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/income?shop_id=%d", t.host, shopID), nil)

	if err != nil {
		return 0, err
	}

	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	log.Printf("%s", string(buf))

	defer resp.Body.Close()

	err = json.Unmarshal(buf, &res)
	if err != nil {
		return 0, err
	}

	return res.Data.NetIncome, nil
}
