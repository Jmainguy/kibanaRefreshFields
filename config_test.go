package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig(t *testing.T) {
	_, err := getConfig()
	assert.Equal(t, err, nil, "Err should be nil")
}
