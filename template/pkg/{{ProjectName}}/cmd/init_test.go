package cmd

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	initConfig()
	assert.Equal(t, "Hello, template", viper.GetString("welcomeMessage"), "welcome message should be equal")
}
