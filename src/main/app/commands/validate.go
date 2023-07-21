package commands

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/sommerfeld-io/semver/services"
	"github.com/spf13/cobra"
)

// Flag to determine if the output should be human readable plain text or JSON
const JSON_FLAG = "json"

func run(cmd *cobra.Command, args []string) {
	version := args[0]
	validationResult := services.IsValid(version)

	if cmd.Flags().Lookup(JSON_FLAG).Changed {
		bytes, err := json.Marshal(validationResult)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(bytes))
	} else {
		fmt.Println(validationResult.Valid)
	}
}

// NewCmdConfig initializes the `tron config` command.
func NewCmdValidate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate given version string, returns true or false",
		Long:  "Validate given version string against semantic versioning rules, returns true if version is valid, false if version does not comply with rules for sommerfeld-io projects",
		Args:  cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			run(cmd, args)
		},
	}

	cmd.Flags().Bool(JSON_FLAG, false, "return result as json with error message if validation fails")

	return cmd
}
