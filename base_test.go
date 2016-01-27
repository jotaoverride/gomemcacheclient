package gomemcacheclient

import (
	"log"
	"os"
	"testing"
	"github.com/mercadolibre/godocker"
)

var memcachedPort int
var c godocker.ContainerID

func TestMain(m *testing.M) {
	log.Println("Init test")
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	c, _, memcachedPort, _ = godocker.SetupMemcachedContainer()
}

func teardown() {
	log.Println("Test finished")
	c.KillRemove()
}
