package ysmrr

import (
	"github.com/chelnak/ysmrr/pkg/colors"
	"github.com/fatih/color"
)

// Spinner manages a single spinner
type spinner struct {
	spinnerColor  *color.Color
	completeColor *color.Color
	errorColor    *color.Color
	msg           string
	complete      bool
	err           bool
}

// Update updates the spinner message
func (s *spinner) Update(msg string) {
	s.msg = msg
}

// Complete marks the spinner as complete
func (s *spinner) Complete() {
	s.complete = true
}

// Error marks the spinner as error
func (s *spinner) Error() {
	s.err = true
}

func NewSpinner(msg string, spinnerColor, completeColor, errorColor colors.Color) *spinner {
	return &spinner{
		spinnerColor:  colors.GetColor(spinnerColor),
		completeColor: colors.GetColor(completeColor),
		errorColor:    colors.GetColor(errorColor),
		msg:           msg,
	}
}
