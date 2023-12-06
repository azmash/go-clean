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

type ShopMembership struct {
	host string
}

func NewHTTPShopMembership(host string) *ShopMembership {
	return &ShopMembership{
		host: host,
	}
}

func (s *ShopMembership) SetMembership(ctx context.Context, shopID int64, newStatus int) error {
	var res model.ClassResult
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/shop/status?shop_id=%d&new_status=%d", s.host, shopID, newStatus), nil)

	if err != nil {
		return err
	}

	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Printf("%s", string(buf))

	defer resp.Body.Close()

	err = json.Unmarshal(buf, &res)
	if err != nil {
		return err
	}

	return nil
}
