# GoMemCacheClient

GoMemCacheClient is the go client library to interact with memcached service

## Usage

Import client:

	import "github.com/mercadolibre/gomemcacheclient"

Register your cluster:

	gomemcacheclient.RegisterMemChache("cluster")

Get:

	gomemcacheclient.Get(key)

Save: 

	gomemcacheclient.Save(key, value)
	
Update:

	gomemcacheclient.Update(key, value)

Delete:

	gomemcacheclient.Delete(key)


## Run tests

	./run-test.sh

## Development

	In development environment client use an in-memory database

## Questions?

Ask juan.alcoleas@mercadolibre.com (GitHub: jotaoverride)