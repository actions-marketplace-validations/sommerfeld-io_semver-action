package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldBeValid(t *testing.T) {
	testCases := []struct {
		version string
	}{
		// Valid releases
		{"v0.0.1"},
		{"v0.1.0"},
		{"v1.0.0"},
		// Valid pre-releases
		{"v0.0.1-alpha"},
		{"v0.0.1-alpha.0"},
		{"v0.0.1-beta"},
		{"v0.0.1-beta.0"},
		{"v0.1.0-alpha"},
		{"v0.1.0-alpha.0"},
		{"v0.1.0-beta"},
		{"v0.1.0-beta.0"},
		{"v1.0.0-alpha"},
		{"v1.0.0-alpha.0"},
		{"v1.0.0-beta"},
		{"v1.0.0-beta.0"},
		// Valid metadata
		{"v0.0.1+metadata"},
		{"v0.1.0+metadata"},
		{"v1.0.0+metadata"},
		{"v0.0.1-alpha+metadata"},
		{"v0.0.1-alpha.0+metadata"},
		{"v0.0.1-beta+metadata"},
		{"v0.0.1-beta.0+metadata"},
	}
	for _, tc := range testCases {
		assert := assert.New(t)

		got := IsValid(tc.version)

		assert.NotNil(got)
		assert.Empty(got.Error)
		assert.True(got.Valid, "Expected version to be valid: "+tc.version)
	}
}

// func Test_ShouldNotBeValid(t *testing.T) {
// 	testCases := []struct {
// 		version string
// 	}{
// 		// Invalid Syntax: version too low
// 		{"v0.0.0"},
// 		{"v0.0.0-alpha"},
// 		{"v0.0.0-beta"},
// 	}
// 	for _, tc := range testCases {
// 		assert := assert.New(t)

// 		got, err := IsValid(tc.version)

// 		assert.Nil(err)
// 		assert.NotNil(got)
// 		assert.False(got)
// 	}
// }

func Test_ShouldCauseError(t *testing.T) {
	testCases := []struct {
		version string
	}{
		// Invalid Syntax: always use all 3 digits (only pre-release is optional)
		{"v1"},
		{"v1.0"},
		// Invalid Syntax: always with a leading "v"
		{"0.0.1"},
		{"0.1.0"},
		{"1.0.0"},
		// Invalid pre-release: only alpha and beta and always lowercase
		{"v1.0.0-SNAPSHOT"},
		{"v1.0.0-snapshot"},
		{"v0.1.0-Alpha"},
		{"v0.1.0-Beta"},
		{"v0.1.0-rc"},
		{"v0.1.0-release-candidate"},
		{"v0.1.0-anything-else"},
		// Invalid separator
		{"0-1-0"},
		{"v0-1-0"},
		{"v0.1.0.alpha"},
		{"v0.1.0.beta"},
		// Invalid
		{"SNAPSHOT"},
		{"main"},
		{"alpha"},
		{"beta"},
		{"anything-else"},
	}
	for _, tc := range testCases {
		assert := assert.New(t)

		got := IsValid(tc.version)

		assert.NotNil(got)
		assert.NotNil(got.Error, "Expected exception because version '"+tc.version+"' should not be valid")
		assert.False(got.Valid)
	}
}
