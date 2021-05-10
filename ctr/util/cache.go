package util

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var c *cache.Cache

func init() {
	if c == nil {
		c = cache.New(5 * time.Minute, 10 * time.Minute)
	}
}

func GetCache() *cache.Cache {
	return c
}
