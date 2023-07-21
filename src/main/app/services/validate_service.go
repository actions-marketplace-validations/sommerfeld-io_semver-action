package services

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Masterminds/semver"
)

func handleNotParsableError(err error) (bool, error) {
	return false, err
}

func hasValidPrefix(version *semver.Version) bool {
	validPrefix := "v"
	return strings.HasPrefix(version.Original(), validPrefix)
}

func hasValidSuffix(version *semver.Version) bool {
	if version.Prerelease() == "" {
		return true
	}

	return version.Prerelease() == "SNAPSHOT"
}

func printErrorMessages(msgs []error) {
	for _, msg := range msgs {
		fmt.Println(msg)
	}
}

// This function validates a provided version string according to the rules of Semantic Versioning.
// Actual validations are done by https://github.com/Masterminds/semver package.
//
// Returns “true“ when the version string is valid and conforms to the Semantic Versioning rules.
func IsValid(version string) (bool, error) {
	v, err := semver.NewVersion(version)
	if err != nil {
		return handleNotParsableError(err)
	}

	isValid := hasValidPrefix(v)
	if !isValid {
		return isValid, errors.New("must start with a leading 'v' (e.g. 'v0.1.0')")
	}

	isValid = hasValidSuffix(v)
	if !isValid {
		return isValid, errors.New("if version has a pre-release suffix, it must equal 'SNAPSHOT'")
	}

	c, err := semver.NewConstraint(">= 0.0.1-SNAPSHOT")
	if err != nil {
		return handleNotParsableError(err)
	}

	isValid, msgs := c.Validate(v)
	printErrorMessages(msgs)

	return isValid, err
}
