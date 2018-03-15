package consulstruct

import "github.com/hashicorp/consul/api"

func New(c *Config) *Decoder {
	return &Decoder{
		config:c,
	}
}

type Config struct {
	Prefix string
	client *api.Client
}

type Decoder struct {
	config *Config
}

func (d *Decoder) Decode(v interface{}) error {
	var err error


	return err
}


