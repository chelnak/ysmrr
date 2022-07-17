// Package ysmrr provides a simple interface for creating and managing
// multiple spinners.
package ysmrr

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/chelnak/ysmrr/pkg/charmap"
	"github.com/chelnak/ysmrr/pkg/colors"
	"github.com/chelnak/ysmrr/pkg/tput"
)

// SpinnerManager manages spinners
type SpinnerManager interface {
	AddSpinner(msg string) *spinner
	Start()
	Stop()
}

type spinnerManager struct {
	Spinners      []*spinner
	chars         []string
	frameDuration time.Duration
	spinnerColor  colors.Color
	completeColor colors.Color
	errorColor    colors.Color
	writer        io.Writer
	done          chan bool
	ticks         *time.Ticker
	pos           int
}

// AddSpinner adds a new spinner to the manager
func (sm *spinnerManager) AddSpinner(msg string) *spinner {
	spinner := NewSpinner(msg, sm.spinnerColor, sm.completeColor, sm.errorColor)
	sm.Spinners = append(sm.Spinners, spinner)
	return spinner
}

// Start signals that all spinners should start
func (sm *spinnerManager) Start() {
	sm.ticks = time.NewTicker(sm.frameDuration)
	go sm.render()
}

// Stop signals that all spinners should complete
func (sm *spinnerManager) Stop() {
	sm.done <- true
	sm.ticks.Stop()
	defer tput.Cnorm()

}

func (sm *spinnerManager) setNextPos() {
	sm.pos += 1
	if sm.pos >= len(sm.chars) {
		sm.pos = 0
	}
}

func (sm *spinnerManager) renderFrame() {
	for _, s := range sm.Spinners {
		if s.complete {
			s.completeColor.Fprint(sm.writer, "\r✓")
			fmt.Fprintf(sm.writer, " %s\n", s.msg)
		} else if s.err {
			s.errorColor.Fprint(sm.writer, "\r✗")
			fmt.Fprintf(sm.writer, " %s\n", s.msg)
		} else {
			s.spinnerColor.Fprintf(sm.writer, "%s", sm.chars[sm.pos])
			fmt.Fprintf(sm.writer, " %s\r", s.msg)
			fmt.Fprint(sm.writer, "\n")
		}
	}
	sm.setNextPos()
}

func (sm *spinnerManager) render() {
	tput.Sc()
	tput.Civis()

	for {
		select {
		case <-sm.done:
			return
		case <-sm.ticks.C:
			sm.renderFrame()
		}

		tput.Rc()
	}
}

// Option represents a spinner manager option.
type Option func(*spinnerManager)

// NewSpinnerManager creates a new spinner manager.
func NewSpinnerManager(options ...Option) SpinnerManager {
	sm := &spinnerManager{
		chars:         charmap.Dots,
		frameDuration: 250 * time.Millisecond,
		spinnerColor:  colors.FgHiGreen,
		errorColor:    colors.FgHiRed,
		completeColor: colors.FgHiGreen,
		writer:        os.Stdout,
		done:          make(chan bool),
	}

	for _, option := range options {
		option(sm)
	}

	return sm
}

// WithChars sets the characters used for the spinners.
func WithCharMap(chars []string) Option {
	return func(sm *spinnerManager) {
		sm.chars = chars
	}
}

// WithFrameDuration sets the duration of each frame.
func WithFrameDuration(d time.Duration) Option {
	return func(sm *spinnerManager) {
		sm.frameDuration = d
	}
}

// WithSpinnerColor sets the color of the spinners.
func WithSpinnerColor(c colors.Color) Option {
	return func(sm *spinnerManager) {
		sm.spinnerColor = c
	}
}

// WithWriter sets the writer used for the spinners.
func WithWriter(w io.Writer) Option {
	return func(sm *spinnerManager) {
		sm.writer = w
	}
}
