package commands

import (
	"bytes"
	"strconv"
	"strings"
	"testing"

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
	assert := assert.New(t)

	var output bytes.Buffer
	cmd := NewCmdValidate()
	args := []string{"v0.1.0"}

	runValidate(&output, cmd, args)

	resultString := strings.ReplaceAll(output.String(), "\n", "")
	result, err := strconv.ParseBool(resultString)

	assert.Nil(err)
	assert.True(result)
}
