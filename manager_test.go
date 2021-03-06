package ysmrr_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/chelnak/ysmrr"
	"github.com/chelnak/ysmrr/pkg/charmap"
	"github.com/chelnak/ysmrr/pkg/colors"
	"github.com/stretchr/testify/assert"
)

func TestNewSpinnerManager(t *testing.T) {
	spinnerManager := ysmrr.NewSpinnerManager()
	assert.NotNil(t, spinnerManager)
}

func TestNewSpinnerManager_WithWriter(t *testing.T) {
	var buf bytes.Buffer
	spinnerManager := ysmrr.NewSpinnerManager(
		ysmrr.WithWriter(&buf),
	)

	assert.Equal(t, &buf, spinnerManager.GetWriter())
}

func TestNewSpinnerManager_WithCharMap(t *testing.T) {
	spinnerManager := ysmrr.NewSpinnerManager(
		ysmrr.WithCharMap(charmap.Arrows),
	)

	assert.Equal(t, charmap.Arrows, spinnerManager.GetCharMap())
}

func TestNewSpinnerManager_WithFrameDuration(t *testing.T) {
	spinnerManager := ysmrr.NewSpinnerManager(
		ysmrr.WithFrameDuration(2 * time.Second),
	)

	assert.Equal(t, 2*time.Second, spinnerManager.GetFrameDuration())
}

func TestAddSpinner(t *testing.T) {
	spinnerManager := ysmrr.NewSpinnerManager()
	spinnerManager.AddSpinner("test")
	assert.Equal(t, 1, len(spinnerManager.GetSpinners()))
}

func TestGetWriter(t *testing.T) {
	var buf bytes.Buffer
	spinnerManager := ysmrr.NewSpinnerManager(
		ysmrr.WithWriter(&buf),
	)

	assert.Equal(t, &buf, spinnerManager.GetWriter())
}

func TestGetCharMap(t *testing.T) {
	spinnerManager := ysmrr.NewSpinnerManager(
		ysmrr.WithCharMap(charmap.Arrows),
	)

	assert.Equal(t, charmap.Arrows, spinnerManager.GetCharMap())
}

func TestGetFrameDuration(t *testing.T) {
	spinnerManager := ysmrr.NewSpinnerManager(
		ysmrr.WithFrameDuration(2 * time.Second),
	)

	assert.Equal(t, 2*time.Second, spinnerManager.GetFrameDuration())
}

func TestGetSpinnerColor(t *testing.T) {
	spinnerManager := ysmrr.NewSpinnerManager(
		ysmrr.WithSpinnerColor(colors.FgHiRed),
	)

	assert.Equal(t, colors.FgHiRed, spinnerManager.GetSpinnerColor())
}

func TestGetCompleteColor(t *testing.T) {
	spinnerManager := ysmrr.NewSpinnerManager(
		ysmrr.WithCompleteColor(colors.FgHiGreen),
	)

	assert.Equal(t, colors.FgHiGreen, spinnerManager.GetCompleteColor())
}

func TestGetErrorColor(t *testing.T) {
	spinnerManager := ysmrr.NewSpinnerManager(
		ysmrr.WithErrorColor(colors.FgHiRed),
	)

	assert.Equal(t, colors.FgHiRed, spinnerManager.GetErrorColor())
}

func TestGetMessageColor(t *testing.T) {
	spinnerManager := ysmrr.NewSpinnerManager(
		ysmrr.WithMessageColor(colors.FgHiBlue),
	)

	assert.Equal(t, colors.FgHiBlue, spinnerManager.GetMessageColor())
}
