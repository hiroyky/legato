package gql

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncodeID(t *testing.T) {
	actual := EncodeID("User", 3)
	expected := "VXNlcjoz"
	assert.Equal(t, expected, actual)
}

func TestEncodeID_異常系(t *testing.T) {
	defer func() {
		err := recover()
		if err == nil {
			t.Errorf("panic dosen't occured")
		}
	}()
	EncodeID("Use:r", 3)
}

func TestDecodeID(t *testing.T) {
	actual, err := DecodeID("VXNlcjoz")
	expected := int64(3)
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}
