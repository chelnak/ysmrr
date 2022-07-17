package ysmrr_test

import (
	"testing"

	"github.com/chelnak/ysmrr"
	"github.com/chelnak/ysmrr/pkg/colors"
	"github.com/stretchr/testify/assert"
)

func TestSpinner(t *testing.T) {
	initialMessage := "test"
	spinnerOptions := ysmrr.SpinnerOptions{
		Message:       initialMessage,
		SpinnerColor:  colors.FgHiGreen,
		CompleteColor: colors.FgHiGreen,
		ErrorColor:    colors.FgHiRed,
		MessageColor:  colors.NoColor,
	}

	spinner := ysmrr.NewSpinner(spinnerOptions)

	assert.Equal(t, initialMessage, spinner.GetMessage())

	updatedMessage := "test2"
	spinner.UpdateMessage(updatedMessage)
	assert.Equal(t, updatedMessage, spinner.GetMessage())

	spinner.Complete()
	assert.Equal(t, true, spinner.IsComplete())

	spinner.Error()
	assert.Equal(t, true, spinner.IsError())
}
