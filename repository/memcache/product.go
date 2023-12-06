package memcache

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func (m *Memcache) GetProductByID(ctx context.Context, id int64) (ProductDetail, error) {
	data, ok := m.mem.Get(fmt.Sprintf("product_%d", id))
	if !ok {
		return ProductDetail{}, errors.New("ProductID is not exist")
	}
	return data.(ProductDetail), nil
}

func (m *Memcache) SetProduct(ctx context.Context, data ProductDetail) error {
	return m.mem.Add(fmt.Sprintf("product_%d", data.ProductID), data, 60*time.Minute)
}
