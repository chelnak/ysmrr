package ysmrr

import (
	"sync"

	"github.com/chelnak/ysmrr/pkg/colors"
	"github.com/fatih/color"
)

// Spinner manages a single spinner
type spinner struct {
	mutex         sync.Mutex
	spinnerColor  *color.Color
	completeColor *color.Color
	errorColor    *color.Color
	msg           string
	complete      bool
	err           bool
}

// GetMessage returns the current spinner message
func (s *spinner) GetMessage() string {
	return s.msg
}

// IsComplete returns true if the spinner is complete
func (s *spinner) IsComplete() bool {
	return s.complete
}

// IsError returns true if the spinner is in error state
func (s *spinner) IsError() bool {
	return s.err
}

// Update updates the spinner message
func (s *spinner) Update(msg string) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.msg = msg
}

// Complete marks the spinner as complete
func (s *spinner) Complete() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.complete = true
}

// Error marks the spinner as error
func (s *spinner) Error() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.err = true
}

// NewSpinner creates a new spinner instance
func NewSpinner(msg string, spinnerColor, completeColor, errorColor colors.Color) *spinner {
	return &spinner{
		spinnerColor:  colors.GetColor(spinnerColor),
		completeColor: colors.GetColor(completeColor),
		errorColor:    colors.GetColor(errorColor),
		msg:           msg,
	}
}
