package api

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	assert.HTTPSuccess(t, Engine.ServeHTTP, "GET", "/ping", nil)
	assert.HTTPBodyContains(t, Engine.ServeHTTP, "GET", "/ping", nil, "pong")
}
