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
		Prefix:    "openheavens/play",
		QueryOpts: nil,
		Store:     client.KV(),
	})

	if err := decoder.Decode(conf); err != nil {
		log.Fatal(err)
	}

	log.Println(conf)
}
