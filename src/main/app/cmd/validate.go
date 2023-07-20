package cmd

import (
	"github.com/spf13/cobra"
)

// NewCmdConfig initializes the `tron config` command.
func NewCmdValidate() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate: Lorem ipsum ...",       // TODO ... ... ...
		Long:  "Validate: Lorem ipsum dolor ...", // TODO ... ... ...
		Args:  cobra.ExactArgs(0),
	}

	return cmd
}
