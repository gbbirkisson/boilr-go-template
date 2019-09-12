package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v1"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manipulate configuration",
}

var configPrintCmd = &cobra.Command{
	Use:   "print",
	Short: "Prints out effective configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		c := viper.AllSettings()
		bs, err := yaml.Marshal(c)
		if err != nil {
			return err
		}
		fmt.Println(string(bs))
		return nil
	},
}

func init() {
	configCmd.AddCommand(configPrintCmd)
	rootCmd.AddCommand(configCmd)
}
