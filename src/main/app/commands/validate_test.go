package commands

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
	"testing"

	"github.com/sommerfeld-io/semver/services"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_ShouldCreateCmdValidate(t *testing.T) {
	assert := assert.New(t)

	expected := &cobra.Command{
		Use:    "validate",
		Args:   cobra.ExactArgs(1),
		Hidden: false,
	}

	got := NewCmdValidate()

	assert.NotNil(got)
	assert.Equal(expected.Use, got.Use)
	assert.Equal(expected.Hidden, got.Hidden)
	assert.NotEmpty(got.Short)
	assert.NotEmpty(got.Long)
	assert.True(got.Runnable(), "Command should be runnable")
	assert.True(got.IsAvailableCommand(), "Command should be runnable")

	assert.True(got.HasFlags(), "Should have flags")
	assert.True(got.HasAvailableFlags(), "Should have flags, that are not hidden or deprecated")
	assert.True(got.Flags().HasFlags(), "Should have flags")
	assert.True(got.Flags().HasAvailableFlags(), "Should have flags, that are not hidden")

	assert.NotNil(got.Flags().Lookup(JSON_FLAG), "Flag should exist")
	assert.False(got.Flags().Lookup(JSON_FLAG).Hidden, "Flag should not be hidden")
}

func Test_ShouldPrintPlainResult(t *testing.T) {
	testCases := []struct {
		args          []string
		shouldBeValid bool
	}{
		{[]string{"v0.1.0"}, true},
		{[]string{"0.1.0"}, false},
	}
	for _, tc := range testCases {
		assert := assert.New(t)

		var output bytes.Buffer
		valid := runValidate(&output, tc.args, false)

		resultString := strings.ReplaceAll(output.String(), "\n", "")
		result, err := strconv.ParseBool(resultString)

		assert.Nil(err)
		assert.Equal(tc.shouldBeValid, valid, "Version "+tc.args[0])
		assert.Equal(tc.shouldBeValid, result, "Version "+tc.args[0])
	}
}

func Test_ShouldPrintJsonResult(t *testing.T) {
	testCases := []struct {
		args          []string
		shouldBeValid bool
	}{
		{[]string{"v0.1.0"}, true},
		{[]string{"0.1.0"}, false},
	}
	for _, tc := range testCases {
		assert := assert.New(t)

		var output bytes.Buffer
		valid := runValidate(&output, tc.args, true)

		assert.NotNil(output, "Should get json result")

		result := &services.ValidationResult{}
		err := json.Unmarshal(output.Bytes(), result)

		assert.Nil(err)
		assert.Equal(tc.shouldBeValid, valid, "Version "+tc.args[0])
		assert.Equal(tc.args[0], result.Version, "Versions should be the same")
		assert.Equal(tc.shouldBeValid, result.Valid, "Version "+tc.args[0])

		if tc.shouldBeValid {
			assert.Empty(result.Error, "Version should be valid, so error message should be empty")
		} else {
			assert.NotEmpty(result.Error, "Version should NOT be valid, so error message should NOT be empty")
		}
	}
}
