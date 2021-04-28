package service

import (
	"bytes"
	"github.com/dgraph-io/ristretto"
	"github.com/eikona-org/eikona/v2/helper"
)

type CacheService interface {
	CheckCache(key string) (CacheEntry, bool)
	AddToCache(key string, imgWrapper *helper.ImageWrapper)
}

type cacheService struct {
	imageCache *ristretto.Cache
}

type CacheEntry struct {
	ImageType    string
	EncodedImage bytes.Buffer
}

func NewCacheService() CacheService {
	imgCache, err := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.
	})

	if err != nil {
		panic(err)
	}

	return &cacheService{
		imageCache: imgCache,
	}
}

func (service *cacheService) CheckCache(key string) (CacheEntry, bool) {
	cacheVal, found := service.imageCache.Get(key)

	if !found {
		return CacheEntry{}, found
	}

	value, ok := cacheVal.(CacheEntry)

	if !ok {
		panic("Problem occurred while accessing the cache")
	}

	return value, found
}

func (service *cacheService) AddToCache(key string, imgWrapper *helper.ImageWrapper) {
	service.imageCache.Set(key, CacheEntry{ImageType: imgWrapper.ImageType, EncodedImage: *imgWrapper.EncodedImage}, 0)
}
