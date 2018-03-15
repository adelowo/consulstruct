package main

import (
	"log"

	"github.com/adelowo/consulstruct"
	"github.com/hashicorp/consul/api"
)

func main() {

	type config struct {
		Mysql              string   `consul:"x/play/database/mysql"`
		Mongo              string   `consul:"x/play/database/mongo"`
		IsActive           bool     `consul:"x/play/boolean"`
		Count              int      `consul:"x/play/counter"`
		MemcachedEndPoints []string `consul:"x/play/memcached/endpoint" consulSeparator:","`
	}

	conf := new(config)

	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatalf("an error occurred while trying to instantiate client... %v", err)
	}

	decoder := consulstruct.New(&consulstruct.Config{
		Prefix:    "x/play",
		QueryOpts: nil,
		Store:     client.KV(),
	})

	if err := decoder.Decode(conf); err != nil {
		log.Fatal("An error occurred while decoding to struct... %v", err)
	}

	log.Println("Parsed successfully", conf)
}
