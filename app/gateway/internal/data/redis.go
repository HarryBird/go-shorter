package data

import (
	"fmt"
	"time"
)

const (
	cacheZone = "url-shorten"
)

var RedisKeyShortenCodeToURL = cacheEntry{key: "code:%s:url", ttl: 5 * time.Minute}

type cacheEntry struct {
	key string
	ttl time.Duration
}

func (c cacheEntry) extract(v ...interface{}) (string, time.Duration) {
	format := fmt.Sprintf("%s::%s", cacheZone, c.key)
	return fmt.Sprintf(format, v...), c.ttl
}
