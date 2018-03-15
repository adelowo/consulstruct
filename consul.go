package consulstruct

import (
	"errors"
	"reflect"

	"github.com/hashicorp/consul/api"
)

var (
	ErrNonPtr    = errors.New("target struct is not a pointer")
	ErrNotStruct = errors.New("target must be a struct")
)

func New(c *Config) *Decoder {
	return &Decoder{
		config: c,
	}
}

type Config struct {
	Prefix    string
	store     *api.KV
	QueryOpts *api.QueryOptions
}

type Decoder struct {
	config *Config
}

func (d *Decoder) Decode(v interface{}) error {

	pointer := reflect.ValueOf(v)
	if pointer.Kind() != reflect.Ptr {
		return ErrNonPtr
	}

	if pointer.Elem().Kind() != reflect.Struct {
		return ErrNotStruct
	}

	_, _, err := d.config.store.List(d.config.Prefix, d.config.QueryOpts)
	if err != nil {
		return err
	}

	return nil
}
