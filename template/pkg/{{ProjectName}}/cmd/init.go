package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var configFileName string

func setConfigDefaults() {
	viper.SetDefault("loglevel", "debug")
	viper.SetDefault("welcomeMessage", "Hello, template")
}

func initConfig() {
	setConfigDefaults()

	viper.SetConfigName(configFileName)

	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/{{ProjectName}}")

	// Try adding the home directory
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to get home directory: %v\n", err)
	} else {
		viper.AddConfigPath(path.Join(home, ".{{ProjectName}}"))
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error
			fmt.Fprintf(os.Stderr, "%v\n", err)
		} else {
			// Config file was found but another error was produced
			log.Panicf("unable to read config file: %v", err)
		}
	} else {
		// Config file was found and parsed
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			log.Debug("config file changed: " + e.Name)
		})
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// viper.SetEnvPrefix("{{ProjectName}}")
}

func initLogging() {
	ll, err := log.ParseLevel(viper.GetString("loglevel"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to parse loglevel: %v\n", err)
		ll = log.InfoLevel
		fmt.Fprintf(os.Stderr, "using loglevel: %v\n", "info")
	}
	log.SetLevel(ll)
}
