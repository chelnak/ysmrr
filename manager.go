// Package ysmrr provides a simple interface for creating and managing
// multiple spinners.
package ysmrr

import (
	"io"
	"os"
	"os/signal"
	"runtime"
	"time"

	"github.com/chelnak/ysmrr/pkg/charmap"
	"github.com/chelnak/ysmrr/pkg/colors"
	"github.com/chelnak/ysmrr/pkg/tput"
	"github.com/mattn/go-colorable"
)

// SpinnerManager manages spinners
type SpinnerManager interface {
	AddSpinner(msg string) *Spinner
	GetSpinners() []*Spinner
	GetWriter() io.Writer
	GetCharMap() []string
	GetFrameDuration() time.Duration
	GetSpinnerColor() colors.Color
	GetErrorColor() colors.Color
	GetCompleteColor() colors.Color
	GetMessageColor() colors.Color
	Start()
	Stop()
}

type spinnerManager struct {
	spinners      []*Spinner
	chars         []string
	frameDuration time.Duration
	spinnerColor  colors.Color
	completeColor colors.Color
	errorColor    colors.Color
	messageColor  colors.Color
	writer        io.Writer
	done          chan bool
	ticks         *time.Ticker
	frame         int
}

// AddSpinner adds a new spinner to the manager.
func (sm *spinnerManager) AddSpinner(message string) *Spinner {
	opts := SpinnerOptions{
		Message:       message,
		SpinnerColor:  sm.spinnerColor,
		CompleteColor: sm.completeColor,
		ErrorColor:    sm.errorColor,
		MessageColor:  sm.messageColor,
	}
	spinner := NewSpinner(opts)
	sm.spinners = append(sm.spinners, spinner)
	return spinner
}

// GetSpinners returns the spinners managed by the manager.
func (sm *spinnerManager) GetSpinners() []*Spinner {
	return sm.spinners
}

// Start signals that all spinners should start.
func (sm *spinnerManager) Start() {
	// Handle SIGINT and SIGTERM so we can ensure that the
	// terminal is properly reset.
	// Unsure if this is the right place for this especially given
	// that it calls os.Exit.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	go func() {
		<-signals
		sm.Stop()
		os.Exit(0)
	}()

	sm.ticks = time.NewTicker(sm.frameDuration)
	go sm.render()
}

// Stop sends a signal to the render goroutine to stop
// rendering. We then stop the ticker and persist the final
// frame for each spinner.
// Finally the deferred tput command will ensure tat the cursor
// is no longer hidden.
func (sm *spinnerManager) Stop() {
	sm.done <- true
	sm.ticks.Stop()
	defer tput.Cnorm(sm.writer)

	// Persist the final frame for each spinner.
	for _, s := range sm.spinners {
		s.Print(sm.writer, sm.chars[sm.frame])
	}
}

// GetWriter returns the configured io.Writer.
func (sm *spinnerManager) GetWriter() io.Writer {
	return sm.writer
}

// GetCharMap returns the configured character map.
func (sm *spinnerManager) GetCharMap() []string {
	return sm.chars
}

// GetFrameDuration returns the configured frame duration.
func (sm *spinnerManager) GetFrameDuration() time.Duration {
	return sm.frameDuration
}

// GetSpinnerColor returns the configured color of the spinners.
func (sm *spinnerManager) GetSpinnerColor() colors.Color {
	return sm.spinnerColor
}

// GetErrorColor returns the configured color of error icon.
func (sm *spinnerManager) GetErrorColor() colors.Color {
	return sm.errorColor
}

// GetCompleteColor returns the configured color of completed icon.
func (sm *spinnerManager) GetCompleteColor() colors.Color {
	return sm.completeColor
}

// GetMessageColor returns the color of the message.
func (sm *spinnerManager) GetMessageColor() colors.Color {
	return sm.messageColor
}

// This is the code that actually renders the spinners.
// Rendering is done in a separate goroutine so that the main
// goroutine can continue to handle signals.
// The render goroutine is called by Start().
//
// Each tick signal calls renderFrame which in turn will print the current
// frame to the writer provided by the manager.
//
// The render method also emits tput strings to the terminal to set the
// correct location of the cursor.
func (sm *spinnerManager) render() {
	// Prepare the screen.
	tput.Civis(sm.writer)
	tput.BufScreen(sm.writer, len(sm.spinners))
	tput.Cuu(sm.writer, len(sm.spinners))
	tput.Sc(sm.writer)

	for {
		select {
		case <-sm.done:
			return
		case <-sm.ticks.C:
			sm.renderFrame()
		}

		tput.Rc(sm.writer)
	}
}

func (sm *spinnerManager) renderFrame() {
	for _, s := range sm.spinners {
		s.Print(sm.writer, sm.chars[sm.frame])
	}
	sm.setNextFrame()
}

func (sm *spinnerManager) setNextFrame() {
	sm.frame += 1
	if sm.frame >= len(sm.chars) {
		sm.frame = 0
	}
}

// NewSpinnerManager is the constructor for the SpinnerManager.
// You can create a new manager with sensible defaults or you can
// pass in your own options using the provided methods.
//
// For example, this will initialize a default manager:
//
//  sm := NewSpinnerManager()
//
// Or this will initialize a manager with a custom character map:
//
// 	sm := NewSpinnerManager(
// 		WithCharMap(charmap.Arrows)
// 	)
//
// You can pass in multiple options to the constructor:
//
// 	sm := NewSpinnerManager(
// 		WithCharMap(charmap.Arrows),
// 		WithFrameDuration(time.Millisecond * 100),
// 		WithSpinnerColor(colors.Red),
// 	)
func NewSpinnerManager(options ...managerOption) SpinnerManager {
	sm := &spinnerManager{
		chars:         charmap.Dots,
		frameDuration: 100 * time.Millisecond,
		spinnerColor:  colors.FgHiGreen,
		errorColor:    colors.FgHiRed,
		completeColor: colors.FgHiGreen,
		messageColor:  colors.NoColor,
		writer:        getWriter(),
		done:          make(chan bool),
	}

	for _, option := range options {
		option(sm)
	}

	return sm
}

func getWriter() io.Writer {
	// Windows support conveniently provided by github.com/mattn/go-colorable <3.
	if runtime.GOOS == "windows" {
		return colorable.NewColorableStdout()
	} else {
		return os.Stdout
	}
}

// Option represents a spinner manager option.
type managerOption func(*spinnerManager)

// WithCharMap sets the characters used for the spinners.
// Available charmaps can be found in the package github.com/chelnak/ysmrr/pkg/charmap.
// The default charmap is the Dots.
func WithCharMap(chars []string) managerOption {
	return func(sm *spinnerManager) {
		sm.chars = chars
	}
}

// WithFrameDuration sets the duration of each frame.
// The default duration is 250 milliseconds.
func WithFrameDuration(d time.Duration) managerOption {
	return func(sm *spinnerManager) {
		sm.frameDuration = d
	}
}

// WithSpinnerColor sets the color of the spinners.
// Available colors can be found in the package github.com/chelnak/ysmrr/pkg/colors.
// The default color is FgHiGreen.
func WithSpinnerColor(c colors.Color) managerOption {
	return func(sm *spinnerManager) {
		sm.spinnerColor = c
	}
}

// WithErrorColor sets the color of the error icon.
// Available colors can be found in the package github.com/chelnak/ysmrr/pkg/colors.
// The default color is FgHiRed.
func WithErrorColor(c colors.Color) managerOption {
	return func(sm *spinnerManager) {
		sm.errorColor = c
	}
}

// WithCompleteColor sets the color of the complete icon.
// Available colors can be found in the package github.com/chelnak/ysmrr/pkg/colors.
// The default color is FgHiGreen.
func WithCompleteColor(c colors.Color) managerOption {
	return func(sm *spinnerManager) {
		sm.completeColor = c
	}
}

// WithMessageColor sets the color of the message.
// Available colors can be found in the package github.com/chelnak/ysmrr/pkg/colors.
// The default color is NoColor.
func WithMessageColor(c colors.Color) managerOption {
	return func(sm *spinnerManager) {
		sm.messageColor = c
	}
}

// WithWriter sets the writer used for the spinners.
// The writer can be anything that implements the io.Writer interface.
// The default writer is os.Stdout.
func WithWriter(w io.Writer) managerOption {
	return func(sm *spinnerManager) {
		sm.writer = w
	}
}
