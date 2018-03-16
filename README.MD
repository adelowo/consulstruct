### Consulstruct

Decode values from consul into a struct.


Supported types :

- [ ] `[]string`
- [ ] string
- [ ] int
- [ ] bool

```go

package main

import (
	"log"

	"github.com/adelowo/consulstruct"
	"github.com/hashicorp/consul/api"
)

func main() {

	type config struct {
		MemcachedEndPoints []string `consul:"memcached/endpoint" consulSeparator:","`
		Mysql              string   `consul:"database/mysql"`
		Mongo              string   `consul:"database/mongo"`
		IsActive           bool     `consul:"boolean"`
		Count              int      `consul:"counter"`
	}

	conf := new(config)

	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatalf("an error occurred while trying to instantiate client... %v", err)
	}

	decoder := consulstruct.New(&consulstruct.Config{
		// This prefix would be applied to all elements in the config struct
		// For example -> database/mysql would be evaluated as x/play/database/mysql
		Prefix:    "x/play",
		QueryOpts: nil,
		Store:     client.KV(),
	})

	if err := decoder.Decode(conf); err != nil {
		log.Fatal(err)
	}

	log.Println(conf)
}

```

