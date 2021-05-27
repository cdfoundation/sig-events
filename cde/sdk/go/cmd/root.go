package cmd

import (
	"fmt"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"

)

var (
	// Used for flags.
	cfgFile     string

	rootCmd = &cobra.Command{
		Use:   "cde",
		Short: "CloudEvents Emitter for Continuous Delivery Events",
		Long: `CD - Continuous Delivery is a simple library and CLI to emit CloudEvents related to Continuous Delivery.`,
	}
)

var CDE_SINK = os.Getenv("CDE_SINK")
var source = "cde-cli"

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}


func init() {

	if CDE_SINK == "" {
		CDE_SINK = "http://localhost:8080"
	}

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cde.yaml)")


}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}


func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, _ := homedir.Dir()
		//cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

