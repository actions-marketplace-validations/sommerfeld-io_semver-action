package commands

import (
	"bytes"
	"io"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_ShouldCreateCmdRoot(t *testing.T) {
	expected := &cobra.Command{
		Use:  "semver",
		Args: cobra.ExactArgs(0), // Todo: validate Args count
	}

	got := NewCmdRoot()
	assert.NotNil(t, got)
	assert.Equal(t, expected.Use, got.Use)
	assert.NotEmpty(t, got.Short)
	assert.NotEmpty(t, got.Long)
	assert.False(t, got.Runnable(), "Command should NOT be runnable")
	assert.False(t, got.IsAvailableCommand(), "Command should be runnable")
	// Todo: validate Args count
}

func Test_ShouldPrintHelpText(t *testing.T) {
	got := NewCmdRoot()
	outputStream := bytes.NewBufferString("")
	got.SetOut(outputStream)
	got.SetErr(outputStream)

	assert.NotNil(t, got)

	err := got.Execute()
	assert.Nil(t, err)
	if err != nil {
		t.Fatal(err)
	}

	out, err := io.ReadAll(outputStream)
	assert.Nil(t, err)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, string(out))
	assert.Contains(t, string(out), got.Long)
}
