package memcache

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type ProductDetail struct {
	ProductID     int64
	ProductName   string
	ProductStatus int
}

type Memcache struct {
	mem *cache.Cache
}

func NewMemcache() *Memcache {
	c := cache.New(5*time.Minute, 10*time.Minute)
	return &Memcache{
		mem: c,
	}
}
