package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShouldBeValid(t *testing.T) {
	testCases := []struct {
		version string
	}{
		{"v0.0.1"},
		{"v0.0.1-SNAPSHOT"},
		{"v0.1.0"},
		{"v0.1.0-SNAPSHOT"},
		{"v1.0.0"},
		{"v1.0.0-SNAPSHOT"},
	}
	for _, tc := range testCases {
		assert := assert.New(t)

		got, err := IsValid(tc.version)

		assert.Nil(err)
		assert.NotNil(got)
		assert.True(got, "Version '"+tc.version+"' should ve valid")
	}
}

func Test_ShouldNotBeValid(t *testing.T) {
	testCases := []struct {
		version string
	}{
		{"0.0.1"},
		{"0.0.1-SNAPSHOT"},
		{"0.0.1-snapshot"},
		{"0.1.0"},
		{"0.1.0-SNAPSHOT"},
		{"0.1.0-snapshot"},
		{"1.0.0"},
		{"1.0.0-SNAPSHOT"},
		{"1.0.0-snapshot"},
		{"0.1.0-alpha"},
		{"0.1.0-beta"},
		{"0.1.0-rc"},
		{"0.1.0-release-candidate"},
		{"0.1.0-anything-else"},
	}
	for _, tc := range testCases {
		assert := assert.New(t)

		got, err := IsValid(tc.version)

		assert.NotNil(err)
		assert.NotNil(got)
		assert.False(got, "Version '"+tc.version+"' should be invalid")
	}
}

func Test_ShouldNotBeParsable(t *testing.T) {
	testCases := []struct {
		version string
	}{
		{"0-1-0"},
		{"0-1-0-SNAPSHOT"},
		{"v0-1-0"},
		{"v0-1-0-SNAPSHOT"},
		{"SNAPSHOT"},
		{"main"},
		{"alpha"},
		{"beta"},
		{"anything-else"},
	}
	for _, tc := range testCases {
		assert := assert.New(t)

		_, err := IsValid(tc.version)

		assert.NotEmpty(err.Error())
	}
}
