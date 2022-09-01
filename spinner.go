package ysmrr

import (
	"fmt"
	"io"
	"sync"

	"github.com/chelnak/ysmrr/pkg/colors"
	"github.com/fatih/color"
)

// Spinner manages a single spinner
type Spinner struct {
	mutex         sync.Mutex
	spinnerColor  *color.Color
	completeColor *color.Color
	errorColor    *color.Color
	messageColor  *color.Color
	message       string
	complete      bool
	err           bool
	hasUpdate     chan bool
}

// GetMessage returns the current spinner message.
func (s *Spinner) GetMessage() string {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.message
}

// UpdateMessage updates the spinner message.
func (s *Spinner) UpdateMessage(message string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.message = message
	s.notifyHasUpdate()
}

// IsComplete returns true if the spinner is complete.
func (s *Spinner) IsComplete() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.complete
}

// IsError returns true if the spinner is in error state.
func (s *Spinner) IsError() bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.err
}

// Complete marks the spinner as complete.
func (s *Spinner) Complete() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.complete = true
}

// Error marks the spinner as error.
func (s *Spinner) Error() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.err = true
}

// Print prints the spinner at a given position.
func (s *Spinner) Print(w io.Writer, char string) {
	if s.IsComplete() {
		print(w, "✓", s.completeColor)
	} else if s.IsError() {
		print(w, "✗", s.errorColor)
	} else {
		print(w, char, s.spinnerColor)
	}

	message := fmt.Sprintf(" %s\r\n", s.message)
	print(w, message, s.messageColor)
}

func print(w io.Writer, s string, c *color.Color) {
	if c != nil {
		_, _ = c.Fprintf(w, s)
	} else {
		fmt.Fprint(w, s)
	}
}

func (s *Spinner) notifyHasUpdate() {
	select {
	case s.hasUpdate <- true:
	default:
	}
}

type SpinnerOptions struct {
	SpinnerColor  colors.Color
	CompleteColor colors.Color
	ErrorColor    colors.Color
	MessageColor  colors.Color
	Message       string
	HasUpdate     chan bool
}

// NewSpinner creates a new spinner instance.
func NewSpinner(options SpinnerOptions) *Spinner {
	return &Spinner{
		spinnerColor:  colors.GetColor(options.SpinnerColor),
		completeColor: colors.GetColor(options.CompleteColor),
		errorColor:    colors.GetColor(options.ErrorColor),
		messageColor:  colors.GetColor(options.MessageColor),
		message:       options.Message,
		hasUpdate:     options.HasUpdate,
	}
}
