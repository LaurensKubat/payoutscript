package jsonrpc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_NewClient(t *testing.T) {
	assert.NotPanics(t, func() {
		_, err := newClient("0.0.0.0:8080")
		if err != nil {
			panic(err)
		}
	})
}
