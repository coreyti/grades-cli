package cli

import (
	"fmt"
	"os"

	"github.com/coreyti/grades-cli/commands/generate"
	"github.com/coreyti/grades-cli/commands/version"
	"github.com/cppforlife/go-cli-ui/ui"
	"github.com/spf13/cobra"
)

var config string
var command = NewCLI()

// NewCLI returns a new "`grades` command"
func NewCLI() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "grades",
		Short: "Jon's Grading CLI.",
		Long:  `A tool for generating grades from test scores.`,
	}

	ui := ui.NewConfUI(ui.NewNoopLogger())

	cmd.AddCommand(generate.NewCommand(generate.NewOptions(ui)))
	cmd.AddCommand(version.NewCommand(version.NewOptions(ui)))

	return cmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the command.
func Execute() {
	if err := command.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// func init() {
// 	cobra.OnInitialize(initConfig)
//
// 	// Here you will define your flags and configuration settings.
// 	// Cobra supports persistent flags, which, if defined here,
// 	// will be global for your application.
// 	command.PersistentFlags().StringVar(&config, "config", "", "config file (default is $HOME/.config/pws.yaml)")
//
// 	// Cobra also supports local flags, which will only run
// 	// when this action is called directly.
// 	// command.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }
//
// // initConfig reads in config file and ENV variables if set.
// func initConfig() {
// 	if config != "" {
// 		// Use config file from the flag.
// 		viper.SetConfigFile(config)
// 	} else {
// 		// Find home directory.
// 		home, err := homedir.Dir()
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}
//
// 		// Search config in home directory with name ".config/pws" (without extension).
// 		viper.AddConfigPath(home)
// 		viper.SetConfigName(".config/pws")
// 	}
//
// 	viper.AutomaticEnv() // read in environment variables that match
//
// 	// If a config file is found, read it in.
// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Println("Using config file:", viper.ConfigFileUsed())
// 	}
// }
