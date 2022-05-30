package ifttt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	apiKey = "INSERT_KEY_HERE"
)

func TestMissingApiKey(t *testing.T) {
	_, err := NewClient("")
	assert.Error(t, err)
}

func TestTriggerJSONWithKey(t *testing.T) {
	c, err := NewClient(apiKey)
	assert.NoError(t, err)

	err = c.TriggerJSONWithKey("wand_interact", nil)
	assert.NoError(t, err)
}
