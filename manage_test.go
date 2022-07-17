package ysmrr_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/chelnak/ysmrr"
	"github.com/chelnak/ysmrr/pkg/charmap"
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
