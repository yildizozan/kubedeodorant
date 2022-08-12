package cmd

import (
	"fmt"
	"kubed/cmd/greetings"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
)

var flagVerbose bool

var cmdRoot = &cobra.Command{
	Use:   "kubed",
	Short: "Kubedeodorant",
	Run: func(cmd *cobra.Command, args []string) {
		greetings.Hi()
	},
}

var cmdHelp = &cobra.Command{
	Use:   "help",
	Short: "Help",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Help")
	},
}

func Execute() error {
	cmdRoot.PersistentFlags().BoolVarP(&flagVerbose, "verbose", "v", false, "verbose output")
	cmdRoot.AddCommand(cmdHelp)

	cmdRoot.MarkPersistentFlagRequired("port")
	return cmdRoot.Execute()
}

func init() {
	// Flags
	//cmdRoot.PersistentFlags().Int16P("port", "p", 8080, "listening port")

	viper.Set("Verbose", true)

	// Set the file name of the configurations file
	viper.SetConfigName("Kubed")

	// Set the path to look for the configurations file
	viper.AddConfigPath("$HOME/.kuberc")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yaml")

	// Set flags
	//viper.BindPFlag("port", cmdRoot.PersistentFlags().Lookup("port"))

	// Set undefined variables
	viper.SetDefault("version", "0.0.1")

	// If config file exists
	//if err := viper.ReadInConfig(); err != nil {
	//	fmt.Println("fatal error config file: default \n", err)
	//	os.Exit(1)
	//}

}