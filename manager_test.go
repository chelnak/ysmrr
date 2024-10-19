package ysmrr_test

import (
	"bytes"
	"sync"
	"testing"
	"time"

	"github.com/chelnak/ysmrr"
	"github.com/chelnak/ysmrr/pkg/animations"
	"github.com/chelnak/ysmrr/pkg/colors"
	"github.com/stretchr/testify/assert"
)

var _, arrows = animations.GetAnimation(animations.Arrow)

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

func TestNewSpinnerManager_WithAnimation(t *testing.T) {
	spinnerManager := ysmrr.NewSpinnerManager(
		ysmrr.WithAnimation(animations.Arrow),
	)

	assert.Equal(t, arrows, spinnerManager.GetAnimation())
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

func TestGetAnimation(t *testing.T) {
	spinnerManager := ysmrr.NewSpinnerManager(
		ysmrr.WithAnimation(animations.Arrow),
	)

	assert.Equal(t, arrows, spinnerManager.GetAnimation())
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

func TestManagerConcurrentSpinnerUsage(t *testing.T) {
	var out bytes.Buffer
	spinnerManager := ysmrr.NewSpinnerManager(ysmrr.WithWriter(&out))
	spinnerManager.Start()

	var wg sync.WaitGroup
	for _, msg := range []string{"test-1", "test-2", "test-3"} {
		wg.Add(1)
		go func(wg *sync.WaitGroup, msg string) {
			defer wg.Done()

			s := spinnerManager.AddSpinner(msg)
			s.UpdateMessage(msg + "-update")
			s.Complete()
		}(&wg, msg)
	}
	wg.Wait()

	spinnerManager.Stop()
}

func TestManagerRunning(t *testing.T) {
	spinnerManager := ysmrr.NewSpinnerManager()
	assert.False(t, spinnerManager.Running())
	spinnerManager.Start()
	assert.True(t, spinnerManager.Running())
	spinnerManager.Stop()
	assert.False(t, spinnerManager.Running())
}

func TestManagerStopTwice(t *testing.T) {
	spinnerManager := ysmrr.NewSpinnerManager()
	spinnerManager.Start()
	spinnerManager.AddSpinner("test")
	spinnerManager.Stop()
	spinnerManager.Stop()
}
