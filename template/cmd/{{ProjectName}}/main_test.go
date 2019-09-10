package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	assert.Equal(t, "Hello, template", message(), "welcome message should be equal")
}
