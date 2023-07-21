package services

import (
	"errors"
	"strings"

	"github.com/hashicorp/go-version"
)

func handleError(err error) (bool, error) {
	return false, err
}

func validatePrefix(v *version.Version) error {
	if strings.HasPrefix(v.Original(), "v") {
		return nil
	}
	return errors.New("must start with a leading v: " + v.Original())
}

func validatePreRelease(v *version.Version) error {
	if strings.Contains(v.Prerelease(), "alpha") {
		return nil
	}
	if strings.Contains(v.Prerelease(), "beta") {
		return nil
	}
	if v.Prerelease() == "" {
		return nil
	}

	return errors.New("invalid pre-release mark (use alpha or beta): " + v.Original())
}

func validateSyntax(v *version.Version) error {
	dotCount := strings.Count(v.Original(), ".")
	if dotCount == 2 || dotCount == 3 {
		return nil
	}

	return errors.New("invalid format (use major.minor.patch / major.minor.patch-prerelease / major.minor.patch-prerelease.counter): " + v.Original())
}

// This function validates a provided version string according to the rules of Semantic Versioning.
// Actual validations are done by https://github.com/hashicorp/go-version package and by some custom
// checks.
//
// Returns “true“ when the version string is valid and conforms to the Semantic Versioning rules.
//
// == See also
// link:/dev-environment-config/main/development-guidelines.html[Semantic Versioning section in out Development Principles]
func IsValid(_version string) (bool, error) {
	v, err := version.NewVersion(_version)
	if err != nil {
		return handleError(err)
	}

	err = validatePrefix(v)
	if err != nil {
		return handleError(err)
	}

	err = validatePreRelease(v)
	if err != nil {
		return handleError(err)
	}

	err = validateSyntax(v)
	if err != nil {
		return handleError(err)
	}

	c, err := version.NewConstraint(">= 0.0.1-SNAPSHOT")
	if err != nil {
		return handleError(err)
	}

	isValid := c.Check(v)

	return isValid, err
}
