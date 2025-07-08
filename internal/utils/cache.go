package utils

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type URLCache struct {
	cache *cache.Cache
}

func NewURLCache() *URLCache {
	return &URLCache{
		cache: cache.New(5*time.Minute, 10*time.Minute),
	}
}

func (uc *URLCache) Set(key string, url string) {
	uc.cache.Set(key, url, cache.DefaultExpiration)
}

func (uc *URLCache) Get(key string) (string, bool) {
	if cached, found := uc.cache.Get(key); found {
		return cached.(string), true
	}
	return "", false
}
