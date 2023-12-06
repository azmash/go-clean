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

type ShopScore struct {
	host string
}

func NewHTTPShopScore(host string) *ShopScore {
	return &ShopScore{
		host: host,
	}
}

func (s *ShopScore) GetShopScore(ctx context.Context, shopID int64) (int, error) {
	var res model.ShopScoreResult
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/score?shop_id=%d", s.host, shopID), nil)

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

	return res.Data.ShopScore, nil
}
