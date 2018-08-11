package importer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProviders_StringToInt(t *testing.T) {
	i, err := StringToInt("2")
	assert.Equal(t, i, int16(2))
	assert.NoError(t, err)

	i, err = StringToInt("2s")
	assert.Error(t, err)
}

func TestProviders_StringToFloat(t *testing.T) {
	i, err := StringToFloat("2.45")
	assert.Equal(t, i, float64(2.45))
	assert.NoError(t, err)

	i, err = StringToFloat("2s")
	assert.Error(t, err)
}
