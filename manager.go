// Package ysmrr provides a simple interface for creating and managing
// multiple spinners.
package ysmrr

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/chelnak/ysmrr/pkg/tput"
	"github.com/fatih/color"
)

var Dots = []string{
	"⠋",
	"⠙",
	"⠹",
	"⠸",
	"⠼",
	"⠴",
	"⠦",
	"⠧",
	"⠇",
	"⠏",
}

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
	writer        io.Writer
	done          chan bool
	ticks         *time.Ticker
	pos           int
}

// AddSpinner adds a new spinner to the manager
func (sm *spinnerManager) AddSpinner(msg string) *spinner {
	c := color.New(color.FgHiGreen)
	spinner := NewSpinner(msg, c)

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
			fmt.Fprintf(sm.writer, "\r✓ %s\n", s.msg)
		} else if s.err {
			fmt.Fprintf(sm.writer, "\r✗ %s\n", s.msg)
		} else {
			s.c.Fprintf(sm.writer, "%s", sm.chars[sm.pos])
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

// NewSpinnerManager creates a new spinner manager
func NewSpinnerManager() SpinnerManager {
	return &spinnerManager{
		chars:         Dots,
		frameDuration: 250 * time.Millisecond,
		writer:        os.Stdout,
		done:          make(chan bool),
	}
}
