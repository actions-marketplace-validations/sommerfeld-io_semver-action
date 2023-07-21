package commands

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_ShouldCreateCmdValidate(t *testing.T) {
	expected := &cobra.Command{
		Use:  "validate",
		Args: cobra.ExactArgs(1), // Todo: validate Args count
	}

	got := NewCmdValidate()
	assert.NotNil(t, got)
	assert.Equal(t, expected.Use, got.Use)
	assert.NotEmpty(t, got.Short)
	assert.NotEmpty(t, got.Long)
	assert.True(t, got.Runnable(), "Command should be runnable")
	assert.True(t, got.IsAvailableCommand(), "Command should be runnable")
	// Todo: validate Args count
}
