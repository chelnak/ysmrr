package ysmrr_test

import (
	"testing"

	"github.com/chelnak/ysmrr"
	"github.com/chelnak/ysmrr/pkg/colors"
	"github.com/stretchr/testify/assert"
)

func TestSpinner(t *testing.T) {
	initialMessage := "test"
	spinner := ysmrr.NewSpinner(initialMessage, colors.FgHiGreen, colors.FgHiYellow, colors.FgHiRed)

	assert.Equal(t, initialMessage, spinner.GetMessage())

	updatedMessage := "test2"
	spinner.Update(updatedMessage)
	assert.Equal(t, updatedMessage, spinner.GetMessage())

	spinner.Complete()
	assert.Equal(t, true, spinner.IsComplete())

	spinner.Error()
	assert.Equal(t, true, spinner.IsError())
}
