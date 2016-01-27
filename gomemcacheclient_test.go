package gomemcacheclient

/**
 * @author jotaoverride
 */

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"fmt"
)

func TestSet(t *testing.T) {
	server := fmt.Sprint("192.168.99.100:", memcachedPort)

	client, err := NewClient([]string{server}, 15 * 60)
	assert.Nil(t, err)

	foo := map[string]interface{}{"key": "value"}
	empty := map[string]interface{}{}

	err = client.Set("foo", empty)
	assert.Nil(t, err)

	err = client.Set("foo", foo, 1)
	assert.Nil(t, err)

	err = client.Set("emptyStr", "")
	assert.Nil(t, err)

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

func TestMiss(t *testing.T) {
	client, err := NewClient([]string{"192.168.99.100:32768"})
	var miss interface{}

	err = client.Get("foo", &miss)
	assert.Equal(t, errors.New("memcache: cache miss"), err)
	assert.Equal(t, nil, miss)
}
