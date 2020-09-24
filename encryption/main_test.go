package sdk

import (
	"os"
	"testing"
)
import "github.com/stretchr/testify/assert"

func TestDecryptToString(t *testing.T) {
	input := "abc"
	key := "zzz"
	os.Setenv("KEY", key)

	output, err := EncryptToString(input, "KEY")

	assert.Nil(t, err)
	assert.Equal(t, output, "AAAAAAAAAABIcNsKAjKXqw==")

	result, err := DecryptToString(output, "KEY")

	assert.Nil(t, err)
	assert.Equal(t, result, input)
}

func TestDecryptToStringPanic(t *testing.T) {
	input := "abc"
	key := "zzz"
	os.Setenv("KEY", key)
	os.Setenv("KEY1", "yyyy")

	output, err := EncryptToString(input, "KEY")
	assert.Nil(t, err)

	assert.PanicsWithValue(t, "Decrypt to string invalid", func(){ DecryptToString(output, "KEY1") })
}
