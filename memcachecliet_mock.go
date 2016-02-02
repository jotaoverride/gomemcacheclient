package gomemcacheclient

/**
 * @author jotaoverride
 */

import (
	"github.com/streamrail/concurrent-map"
	"time"
	"errors"
)

type MemcacheClientMock struct {
	memcache          cmap.ConcurrentMap
	defaultExpiration int32
}

func NewClientMock(expiration ...int32) *MemcacheClientMock {
	client := new(MemcacheClientMock)
	client.memcache = cmap.New()

	if expiration != nil {
		client.defaultExpiration = expiration[0]
	} else {
		client.defaultExpiration = 15
	}

	return client
}

func (c *MemcacheClientMock) Get(key string, value interface{}) (err error) {
	v, ok := c.memcache.Get(key)

	if ok {
		decode(v.([]byte), value)
	} else {
		err = errors.New(`memcache: cache miss`)
	}

	return
}

func (c *MemcacheClientMock) Set(key string, value interface{}, expiration ...int32) error {
	v, _ := encode(value)
	c.memcache.Set(key, v)

	var expire time.Duration

	if expiration != nil {
		expire = time.Duration(expiration[0])
	} else {
		expire = time.Duration(c.defaultExpiration)
	}

	go func(memcache *MemcacheClientMock, expirationTime time.Duration) {
			time.Sleep(expire * time.Second)
			c.Delete(key)
		}(c, expire)

	return nil
}

func (c *MemcacheClientMock) Delete(key string) error {
	if _, ok := c.memcache.Get(key); ok {
		c.memcache.Remove(key)
	}
	return nil
}

func (c *MemcacheClientMock) SetDefaultExpiration(secs int32) {
	c.defaultExpiration = secs
}

