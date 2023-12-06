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

type Kyc struct {
	host string
}

func NewHTTPKYC(host string) *Kyc {
	return &Kyc{
		host: host,
	}
}

func (k *Kyc) CheckKYC(ctx context.Context, shopID int64) (bool, error) {
	var res model.KYCResult
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/ktp?shop_id=%d", k.host, shopID), nil)

	if err != nil {
		return false, err
	}

	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return false, err
	}

	defer resp.Body.Close()

	log.Printf("%s", string(buf))
	err = json.Unmarshal(buf, &res)
	if err != nil {
		return false, err
	}

	if res.Data.KycStatus == 1 {
		return true, nil
	}

	return false, nil
}
