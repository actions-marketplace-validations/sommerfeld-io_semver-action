package commands

import (
	"github.com/spf13/cobra"
)

// NewCmdRoot initializes the `tron` root command.
func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "semver",
		Version: "Lorem ...", // TODO https://github.com/sommerfeld-io/semver-action/issues/4
		Short:   "Use Semver to validate versions",
		Long:    "Use Semver to validate versions against semantic verioning rules",
		Args:    cobra.ExactArgs(0),
	}

	return cmd
}

var rootCmd *cobra.Command

func init() {
	rootCmd = NewCmdRoot()
	rootCmd.AddCommand(NewCmdValidate())

	rootCmd.CompletionOptions.HiddenDefaultCmd = true
}

// Execute acts as the entrypoint for the command line interface.
func Execute() error {
	return rootCmd.Execute()
}
