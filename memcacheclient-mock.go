package gomemcacheclient

/**
 * @author jotaoverride
 */

import (
	"github.com/streamrail/concurrent-map"
	"github.com/bradfitz/gomemcache/memcache"
	"time"
	"errors"
)

type MemcacheClientMock struct {
	memcache cmap.ConcurrentMap
}

// Create a new Client.
// Config must have the servers array.
// Optional the defaultExpiration.
func ConnectClient(config map[string]interface{}) (*MemcacheClient, error) {

	selector := new(memcache.ServerList)
	var client *MemcacheClient = nil
	var err error

	if servers, ok := config["servers"]; ok {

		err = selector.SetServers(servers.([]string)...)

		if err == nil {

			client = &MemcacheClient{client: memcache.NewFromSelector(selector)}
			client.client.Timeout = 200 * time.Millisecond // Default timeout: 200 ms.

			defaultExpiration := config["defaultExpiration"]

			switch defaultExpiration.(type) {
			case int32:
				client.defaultExpiration = defaultExpiration.(int32)
			case int:
				client.defaultExpiration = int32(defaultExpiration.(int))
			default:
				client.defaultExpiration = 60 * 15 // Default expiration time: 15 minutes.
			}

			// Test the connection...
			testEncoded, _ := encode("Ok")
			err = client.client.Set(&memcache.Item{Key: "GOLANG_TEST", Value: testEncoded})
		}

	} else {
		err = errors.New("Memcached servers missing.")
	}

	return client, err
}

func (c *MemcacheClientMock) Get(key string, value interface{}) error {
	return nil
}

func (c *MemcacheClientMock) Set(key string, value interface{}, expiration ...int32) error {
	return nil
}

func (c *MemcacheClientMock) Delete(key string) error {
	return nil
}

func (c *MemcacheClientMock) SetDefaultExpiration(secs int32) {
	return
}
