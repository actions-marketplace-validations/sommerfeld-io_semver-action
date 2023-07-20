package cmd

import (
	"github.com/spf13/cobra"
)

// NewCmdRoot initializes the `tron` root command.
func NewCmdRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "semver",
		Version: "Lorem ...",                   // TODO ... ... ...
		Short:   "Root: Lorem ipsum ...",       // TODO ... ... ...
		Long:    "Root: Lorem ipsum dolor ...", // TODO ... ... ...
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
