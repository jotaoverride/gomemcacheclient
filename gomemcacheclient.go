package gomemcacheclient

/**
 * @author jotaoverride
 */

import (
	"os"
)

type Client interface {
	Get(key string, value interface{}) error
	Set(key string, value interface{}, expiration ...int32) error
	Delete(key string) error
	SetDefaultExpiration(secs int32)
}

func NewClient(servers []string, defaultExpiration ...int32) (Client, error) {

	var instance Client
	var err error

	if true || os.Getenv("GO_ENVIRONMENT") == "production" {

		RealInstance := new(MemcacheClient)
		err = RealInstance.connectClient(servers, defaultExpiration...)
		instance = RealInstance

	} else {

		mockInstance := new(MemcacheClientMock)
		instance = mockInstance
	}

	return instance, err
}
