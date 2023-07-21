package commands

import (
	"fmt"

	"github.com/sommerfeld-io/semver/services"
	"github.com/spf13/cobra"
)

// NewCmdConfig initializes the `tron config` command.
func NewCmdValidate() *cobra.Command {
	const SHOW_ERRORS_FLAG = "show-errors"

	cmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate given version string",
		Long:  "Validate given version string against semantic versioning rules - true = version is valid, false = version does not comply with rules for sommerfeld-io projects",
		Args:  cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			version := args[0]

			isValid, err := services.IsValid(version)

			if err != nil && cmd.Flags().Lookup(SHOW_ERRORS_FLAG).Changed {
				fmt.Println(err)
			}
			fmt.Println(isValid)
		},
	}

	cmd.Flags().Bool(SHOW_ERRORS_FLAG, false, "print error messages in case a validation fails")

	return cmd
}
