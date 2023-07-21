package commands

import (
	"fmt"
	"log"

	"github.com/sommerfeld-io/semver/services"
	"github.com/spf13/cobra"
)

// Flag to determine if the output should be human readable plain text or JSON
const JSON_FLAG = "json"

func run(cmd *cobra.Command, args []string) {
	version := args[0]

	isValid, err := services.IsValid(version)

	if cmd.Flags().Lookup(JSON_FLAG).Changed {
		// todo ... add "to json function so something like that"
		fmt.Print("todo")
	} else {
		if err != nil {
			log.Print(err)
		}
		fmt.Println(isValid)
	}
}

// NewCmdConfig initializes the `tron config` command.
func NewCmdValidate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate given version string",
		Long:  "Validate given version string against semantic versioning rules - true = version is valid, false = version does not comply with rules for sommerfeld-io projects",
		Args:  cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			run(cmd, args)
		},
	}

	cmd.Flags().Bool(JSON_FLAG, false, "return result as json")

	return cmd
}
