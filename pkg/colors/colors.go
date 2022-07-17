// Package colors provides a collection of color definitions for use with a spinner.
package colors

import "github.com/fatih/color"

// Color represents an item in the color map.
type Color int

const (
	// FgHiGreen is a foreground high intensity green color.
	FgHiGreen Color = iota

	// FgHiYellow is a foreground high intensity yellow color.
	FgHiYellow

	// FgHiBlue is a foreground high intensity blue color.
	FgHiBlue

	// FgHiRed is a foreground high intensity red color.
	FgHiRed
)

var lookup = map[Color]color.Attribute{
	FgHiGreen:  color.FgHiGreen,
	FgHiYellow: color.FgHiYellow,
	FgHiBlue:   color.FgHiBlue,
	FgHiRed:    color.FgHiRed,
}

// GetColor returns a color.Color for the given color.
func GetColor(c Color) *color.Color {
	return color.New(lookup[c])
}
