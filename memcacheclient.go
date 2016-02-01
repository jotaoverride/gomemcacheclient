package gomemcacheclient

/**
 * @author jotaoverride
 */

import (
	"bytes"
	"encoding/gob"
	"errors"
	"github.com/bradfitz/gomemcache/memcache"
	"time"
)

type MemcacheClient struct {
	client            *memcache.Client
	defaultExpiration int32
}

// Create a new Client.
// Config must have the servers array.
// Optional the defaultExpiration.
func (c *MemcacheClient) connectClient(servers []string, defaultExpiration ...int32) (err error) {

	selector := new(memcache.ServerList)
	err = selector.SetServers(servers...)

	if err == nil {

		c.client = memcache.NewFromSelector(selector)
		c.client.Timeout = 200 * time.Millisecond // Default timeout: 200 ms.

		if defaultExpiration != nil {
			c.defaultExpiration = defaultExpiration[0]
		} else {
			c.defaultExpiration = 60 * 15 // Default expiration time: 15 minutes.
		}

		// Test the connection...
		if _, testErr := c.client.Get("GOLANG_TEST"); testErr != nil && testErr.Error() != "memcache: cache miss" {
			err = errors.New("Can't connect to the servers")
		}
	}

	return
}

func (c *MemcacheClient) SetDefaultExpiration(secs int32) {
	c.defaultExpiration = secs
}

func (c *MemcacheClient) Get(key string, value interface{}) (err error) {
	item, err := c.client.Get(key)
	if err == nil {
		err = decode(item.Value, value)
	}
	return err
}

func (c *MemcacheClient) Set(key string, value interface{}, expiration ...int32) (err error) {
	valueEncoded, err := encode(value)
	if err == nil {
		if expiration != nil {
			err = c.client.Set(&memcache.Item{Key: key, Value: valueEncoded, Expiration: expiration[0]})
		} else {
			err = c.client.Set(&memcache.Item{Key: key, Value: valueEncoded, Expiration: c.defaultExpiration})
		}
	}
	return
}

func (c *MemcacheClient) Delete(key string) error {
	return c.client.Delete(key)
}

// Auxiliary functions

func encode(value interface{}) (encodedValue []byte, err error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)

	err = enc.Encode(value)
	if err == nil {
		encodedValue = buf.Bytes()
	}

	return
}

func decode(value []byte, result interface{}) error {
	dec := gob.NewDecoder(bytes.NewBuffer(value))
	return dec.Decode(result)
}
