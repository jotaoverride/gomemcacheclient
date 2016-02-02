package gomemcacheclient

import (
	"github.com/mercadolibre/godocker"
	"log"
	"os"
	"testing"
)

var memcacheAddrs string
var client Client
var memcacheContainer godocker.ContainerID

func TestMain(m *testing.M) {
	log.Println("Init mock test")
	setupMock()
	codeMock := m.Run()
	log.Println("Init docker test")
	setupDocker()
	codeDocker := m.Run()
	teardown()
	os.Exit(codeDocker + codeMock)
}

func setupMock() {
	os.Setenv("GO_ENVIRONMENT", "testing")
	var err error
	client, err = NewClient([]string{})
	if err != nil {
		log.Println("Test setup error. Can't create memcached client.", err)
	}
}

func setupDocker() {
	os.Setenv("GO_ENVIRONMENT", "production")
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
	memcacheAddrs = memcachedIP + ":" + memcachedPort
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
