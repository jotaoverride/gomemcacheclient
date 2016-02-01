package gomemcacheclient

import (
	"log"
	"os"
	"testing"
	"github.com/mercadolibre/godocker"
)

var memcacheAddrs string
var client Client
var memcacheContainer godocker.ContainerID

func TestMain(m *testing.M) {
	log.Println("Init test")
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	memcacheContainer, err := godocker.StartContainer("memcached")
	if err != nil {
		log.Println("Test setup error. Memcached no started.", err)
		return
	}
	memcachedIP, err := memcacheContainer.IP()
	if err != nil {
		log.Println("Test setup error. Can't get docker IP.", err)
		return
	}
	memcachedPort, err := memcacheContainer.GetPort("11211")
	if err != nil {
		log.Println("Test setup error. Can't get port mapping for memcached container.", err)
		return
	}
	memcacheAddrs = memcachedIP+":"+memcachedPort
	client, err = NewClient([]string{memcacheAddrs})
	if err != nil {
		log.Println("Test setup error. Can't create memcached client.", err)
		return
	}
}

func teardown() {
	log.Println("Test finished")
	memcacheContainer.KillRemove()
}
