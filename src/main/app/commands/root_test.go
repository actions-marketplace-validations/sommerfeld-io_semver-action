package commands

import (
	"bytes"
	"io"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_ShouldCreateCmdRoot(t *testing.T) {
	assert := assert.New(t)

	expected := &cobra.Command{
		Use:    "semver",
		Args:   cobra.ExactArgs(0),
		Hidden: false,
	}

	got := NewCmdRoot()

	assert.NotNil(got)
	assert.Equal(expected.Use, got.Use)
	assert.Equal(expected.Hidden, got.Hidden)
	assert.NotEmpty(got.Short)
	assert.NotEmpty(got.Long)
	assert.False(got.Runnable(), "Command should NOT be runnable")
	assert.False(got.IsAvailableCommand(), "Command should be runnable")

	assert.False(got.HasFlags(), "Should NOT have flags")
	assert.False(got.HasAvailableFlags(), "Should NOT have flags, that are not hidden or deprecated")
	assert.False(got.Flags().HasFlags(), "Should NOT have flags")
	assert.False(got.Flags().HasAvailableFlags(), "Should NOT have flags, that are not hidden")

}

func Test_ShouldPrintHelpText(t *testing.T) {
	assert := assert.New(t)

	got := NewCmdRoot()
	outputStream := bytes.NewBufferString("")
	got.SetOut(outputStream)
	got.SetErr(outputStream)

	assert.NotNil(got)

	err := got.Execute()
	assert.Nil(err)
	if err != nil {
		t.Fatal(err)
	}

	out, err := io.ReadAll(outputStream)
	assert.Nil(err)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(string(out))
	assert.Contains(string(out), got.Long)
}
