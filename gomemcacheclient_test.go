package gomemcacheclient

/**
 * @author jotaoverride
 */

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMiss(t *testing.T) {
	var miss interface{}

	err := client.Get("foo", &miss)
	assert.Equal(t, errors.New("memcache: cache miss"), err)
	assert.Equal(t, nil, miss)
}

func TestSetAndHit(t *testing.T) {
	foo := map[string]interface{}{"key": "value"}

	err := client.Set("foo", foo, 1)
	assert.Nil(t, err)

	empty := map[string]interface{}{}

	// Hit
	err = client.Get("foo", &empty)
	assert.Nil(t, err)
	assert.Equal(t, foo, empty)
	assert.Equal(t, foo, empty)

	err = client.Delete("foo")
	assert.Nil(t, err)

	// Miss
	empty = map[string]interface{}{}

	err = client.Get("foo", &empty)
	assert.Equal(t, errors.New("memcache: cache miss"), err)
	assert.Equal(t, map[string]interface{}{}, empty)
}

func TestSetTwiceAndHit(t *testing.T) {
	err := client.Set("foo", "")
	assert.Nil(t, err)

	foo := map[string]interface{}{"key": "value"}

	err = client.Set("foo", foo, 60)
	assert.Nil(t, err)

	empty := map[string]interface{}{}

	err = client.Get("foo", &empty)
	assert.Nil(t, err)
	assert.Equal(t, foo, empty)
	assert.Equal(t, foo, empty)
}

func TestGetDeleteAndMiss(t *testing.T)  {
	foo := map[string]interface{}{"key": "value"}

	err := client.Set("foo", foo, 1)
	assert.Nil(t, err)

	err = client.Delete("foo")
	assert.Nil(t, err)

	// Miss
	empty := map[string]interface{}{}

	err = client.Get("foo", &empty)
	assert.Equal(t, errors.New("memcache: cache miss"), err)
	assert.Equal(t, map[string]interface{}{}, empty)
}

func TestGetExpiredItem(t *testing.T)  {
	client.SetDefaultExpiration(1)

	foo := map[string]interface{}{"key": "value"}

	err := client.Set("foo", foo)
	assert.Nil(t, err)

	time.Sleep(2 * time.Second)

	// Miss
	empty := map[string]interface{}{}

	err = client.Get("foo", &empty)
	assert.Equal(t, errors.New("memcache: cache miss"), err)
	assert.Equal(t, map[string]interface{}{}, empty)
}

func TestNewClientWithExpiration(t *testing.T) {
	otherClient, err := NewClient([]string{memcacheAddrs}, 1)
	assert.Nil(t, err)

	foo := map[string]interface{}{"key": "value"}

	err = otherClient.Set("foo", foo)
	assert.Nil(t, err)

	time.Sleep(3 * time.Second)

	// Miss
	empty := map[string]interface{}{}

	err = otherClient.Get("foo", &empty)
	assert.Equal(t, errors.New("memcache: cache miss"), err)
	assert.Equal(t, map[string]interface{}{}, empty)
}