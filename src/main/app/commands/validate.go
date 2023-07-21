package commands

import (
	"fmt"

	"github.com/sommerfeld-io/semver/services"
	"github.com/spf13/cobra"
)

// NewCmdConfig initializes the `tron config` command.
func NewCmdValidate() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate given version string",
		Long:  "Validate given version string against semantic versioning rules",
		Args:  cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			version := args[0]

			isValid, err := services.IsValid(version)

			if err != nil { // && --show-errors
				fmt.Println(err)
			}
			fmt.Println(isValid)
		},
	}

	return cmd
}
