package consulstruct

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecoder_OnlyPointers(t *testing.T) {
	decoder := New(new(Config))

	err := decoder.Decode(struct{}{})
	require.Equal(t, ErrNonPtr, err)
}

func TestDecoder_OnlyStructs(t *testing.T) {
	decoder := New(new(Config))

	err := decoder.Decode(new(string))
	require.Equal(t, ErrNotStruct, err)
}

func Test_normalize(t *testing.T) {

	var tt = []struct {
		prefix, key, expected string
	}{
		{"first", "oops", "first/oops"},
		{"second/", "oops", "second/oops"},
		{"second/", "/oops", "second/oops"},
	}

	for _, val := range tt {
		if s := normalize(val.prefix, val.key); s != val.expected {
			t.Fatalf("Strings do not match... Got %s .. expected %s", s, val.expected)
		}
	}
}
