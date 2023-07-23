package commands

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/sommerfeld-io/semver/services"
	"github.com/spf13/cobra"
)

// Default io.Writer
var W = os.Stdout

// Flag to determine if the output should be human readable plain text or JSON
const JSON_FLAG = "json"

func addFlags(cmd *cobra.Command) {
	cmd.Flags().Bool(JSON_FLAG, false, "return result as json with error message if validation fails")
}

func runValidate(w io.Writer, args []string, asJson bool) bool {
	version := args[0]
	validationResult := services.Validate(version)

	if asJson {
		bytes, err := json.Marshal(validationResult)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(w, string(bytes))
	} else {
		fmt.Fprintln(w, validationResult.Valid)
	}

	return validationResult.Valid
}

// NewCmdConfig initializes the `tron config` command.
func NewCmdValidate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validate",
		Short: "Validate a given version string",
		Long:  "Validate a given version string against semantic versioning rules, returns true if version is valid and false if version does not comply with rules for sommerfeld-io projects",
		Args:  cobra.ExactArgs(1),

		Run: func(cmd *cobra.Command, args []string) {
			asJson := cmd.Flags().Lookup(JSON_FLAG).Changed
			valid := runValidate(W, args, asJson)

			if !valid {
				os.Exit(1)
			}
		},
	}

	addFlags(cmd)

	return cmd
}
