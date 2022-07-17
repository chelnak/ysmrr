package ysmrr

import "github.com/fatih/color"

// Spinner manages a single spinner
type spinner struct {
	c        *color.Color
	msg      string
	complete bool
	err      bool
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

func NewSpinner(msg string, c *color.Color) *spinner {
	return &spinner{
		c:   c,
		msg: msg,
	}
}
