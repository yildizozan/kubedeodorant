package cmd

import (
	"fmt"
	"kubed/cmd/greetings"

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
