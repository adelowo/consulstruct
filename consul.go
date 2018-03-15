package consulstruct

import (
	"errors"
	"reflect"
	"strconv"
	"strings"

	"github.com/hashicorp/consul/api"
)

const (
	consulMainTag   = "consul"
	consulSeparator = "consulSeparator"
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
	Store     *api.KV
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

	pairs, _, err := d.config.Store.List(d.config.Prefix, d.config.QueryOpts)
	if err != nil {
		return err
	}

	return d.decode(pointer.Elem(), pairsToDict(d.config.Prefix, pairs))
}

func (d *Decoder) decode(ref reflect.Value, pairs map[string]*api.KVPair) error {

	for i := 0; i < ref.Type().NumField(); i++ {
		structField := ref.Type().Field(i)
		val := fetch(structField, consulMainTag)

		if len(strings.TrimSpace(val)) == 0 {
			continue
		}

		pair, ok := pairs[normalize(d.config.Prefix, val)]
		if !ok {
			continue
		}

		if err := set(ref.Field(i), structField, pair); err != nil {
			return err
		}
	}

	return nil
}

func pairsToDict(prefix string, pairs api.KVPairs) map[string]*api.KVPair {
	var p = make(map[string]*api.KVPair, len(pairs))

	for _, pair := range pairs {
		p[pair.Key] = pair
	}

	return p
}

func normalize(prefix, str string) string {
	if !strings.HasSuffix(prefix, "/") {
		prefix = prefix + "/"
	}

	return prefix + strings.TrimPrefix(str, "/")
}

func fetch(val reflect.StructField, s string) string {
	return val.Tag.Get(s)
}

func set(field reflect.Value, refType reflect.StructField, pair *api.KVPair) error {
	var err error

	s := string(pair.Value)

	switch field.Kind() {

	case reflect.Slice:

		separator := fetch(refType, consulSeparator)

		if reflect.TypeOf([]string(nil)) == refType.Type {
			splitted := strings.Split(s, separator)
			field.Set(reflect.ValueOf(splitted))
			return nil
		}

		err = errors.New("consulstruct : unsupported operation")
		break

	case reflect.String:
		field.SetString(s)
		break

	case reflect.Int:
		n, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			return err
		}

		field.SetInt(n)
		break

	case reflect.Bool:
		truthy, err := strconv.ParseBool(s)
		if err != nil {
			return err
		}

		field.SetBool(truthy)
		break

	default:
		err = errors.New("consulstruct : unsupported operation")
	}

	return err
}
