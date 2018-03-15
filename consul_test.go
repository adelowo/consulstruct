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
