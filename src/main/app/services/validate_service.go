package services

import (
	"errors"
	"strings"

	"github.com/hashicorp/go-version"
)

func handleError(err error) (bool, error) {
	return false, err
}

func hasValidPrefix(v *version.Version) error {
	const VALID_PREFIX = "v"

	if !strings.HasPrefix(v.Original(), VALID_PREFIX) {
		return errors.New("must start with a leading 'v' (e.g. 'v0.1.0')")
	}
	return nil
}

// Version is considered valid without a pre-release mark (e.g. `v0.1.0`) or with a correct
// pre-release mark (e.g. `v0.1.0-SNAPSHOT`). The function returns `nil` when version is valid
// or an error when version is invalid.
func hasValidSuffix(v *version.Version) error {
	const VALID_SUFFIX = "SNAPSHOT"

	if v.Prerelease() == "" {
		return nil
	}

	if v.Prerelease() == VALID_SUFFIX {
		return nil
	}

	return errors.New("if version has a pre-release suffix, it must equal 'SNAPSHOT'")
}

// This function validates a provided version string according to the rules of Semantic Versioning.
// Actual validations are done by https://github.com/hashicorp/go-version package and by some custom
// checks.
//
// Returns “true“ when the version string is valid and conforms to the Semantic Versioning rules.
func IsValid(_version string) (bool, error) {
	v, err := version.NewVersion(_version)
	if err != nil {
		return handleError(err)
	}

	err = hasValidPrefix(v)
	if err != nil {
		return handleError(err)
	}

	err = hasValidSuffix(v)
	if err != nil {
		return handleError(err)
	}

	// todo needs 3 digits

	c, err := version.NewConstraint(">= 0.0.1-SNAPSHOT")
	if err != nil {
		return handleError(err)
	}

	isValid := c.Check(v)

	return isValid, err
}
