package commands

import (
	"github.com/spf13/cobra"
)

// NewCmdRoot initializes the `semver` root command.
func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "semver",
		Version: "main",     // todo ...
		Short:   "TODO ...", // todo ...
		Long:    "TODO ...", // todo ...
		Args:    cobra.ExactArgs(0),
	}

	return cmd
}

var rootCmd *cobra.Command

func init() {
	rootCmd = NewCmdRoot()
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
}

// Execute acts as the entrypoint for the command line interface.
func Execute() error {
	return rootCmd.Execute()
}
