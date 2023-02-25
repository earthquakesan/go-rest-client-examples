package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfig(t *testing.T) {
	config := getConfig()
	expectedConfig := Config{
		BaseUrl: "http://127.0.0.1:5000",
	}
	assert.Equal(t, expectedConfig, config)
}
