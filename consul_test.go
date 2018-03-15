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
