package services

import (
	"errors"
	"strings"

	"github.com/hashicorp/go-version"
)

// This object represents the result of a semantic version validation and is returned to the user.
// It encapsulates the outcome of the validation process, including information about whether the
// input version is valid or not, along with potential error messages.
type ValidationResult struct {
	Version string `json:"version"`
	Valid   bool   `json:"valid"`
	Error   string `json:"error"`
}

func handleError(result ValidationResult, err error) ValidationResult {
	result.Valid = false
	result.Error = err.Error()
	return result
}

func validatePrefix(v *version.Version) error {
	if strings.HasPrefix(v.Original(), "v") {
		return nil
	}
	return errors.New("must start with a leading 'v'")
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

	return errors.New("invalid pre-release mark: use 'alpha' or 'beta'")
}

func validateSyntax(v *version.Version) error {
	dotCount := strings.Count(v.Original(), ".")
	if dotCount == 2 || dotCount == 3 {
		return nil
	}

	return errors.New("invalid format: use 'major.minor.patch' / 'major.minor.patch-prerelease' / 'major.minor.patch-prerelease.counter')")
}

// This function validates a provided version string according to the rules of Semantic Versioning.
// Actual validations are done by https://github.com/hashicorp/go-version package and by some custom
// checks.
//
// Returns “true“ when the version string is valid and conforms to the Semantic Versioning rules.
//
// == See also
// link:/dev-environment-config/main/development-guidelines.html[Semantic Versioning section in out Development Principles]
func Validate(v string) ValidationResult {
	result := ValidationResult{
		Version: v,
	}

	parsedVersion, err := version.NewVersion(v)
	if err != nil {
		return handleError(result, err)
	}

	err = validatePrefix(parsedVersion)
	if err != nil {
		return handleError(result, err)
	}

	err = validatePreRelease(parsedVersion)
	if err != nil {
		return handleError(result, err)
	}

	err = validateSyntax(parsedVersion)
	if err != nil {
		return handleError(result, err)
	}

	// constraints, err := version.NewConstraint(">= 0.0.1-alpha")
	// if err != nil {
	// 	return handleError(err)
	// }
	// isValid := constraints.Check(parsedVersion)
	// return isValid, nil

	result.Valid = true
	return result
}
