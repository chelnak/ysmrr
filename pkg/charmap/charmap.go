// Package charmap provides a collection of character maps to be used
// with a spinner.
// Spinners have been borrowed from the following sources:
// * https://wiki.tcl-lang.org/page/Text+Spinner
// * https://stackoverflow.com/questions/2685435/cooler-ascii-spinners
package charmap

type CharMap int

const (
	Arc CharMap = iota + 100
	Arrow
	Baloon
	Baloon2
	Circle
	CircleHalves
	CircleQuarters
	Dots
	Hamburger
	Layer
	Pipe
	Point
	Star
	SquareCorners
)

var lookup = map[CharMap][]string{
	Arc:            {"◜", "◠", "◝", "◞", "◡", "◟"},
	Arrow:          {"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"},
	Baloon:         {".", "o", "O", "@", "*"},
	Baloon2:        {".", "o", "O", "°", "O", "o", "."},
	Circle:         {"◡", "⊙", "◠"},
	CircleHalves:   {"◐", "◓", "◑", "◒"},
	CircleQuarters: {"◴", "◷", "◶", "◵"},
	Dots:           {"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
	Hamburger:      {"☱", "☲", "☴"},
	Layer:          {"-", "=", "≡"},
	Pipe:           {"┤", "┘", "┴", "└", "├", "┌", "┬", "┐"},
	Point:          {"∙∙∙", "●∙∙", "∙●∙", "∙∙●", "∙∙∙"},
	Star:           {"✶", "✸", "✹", "✺", "✹", "✷"},
	SquareCorners:  {"◰", "◳", "◲", "◱"},
}

// GetCharMap retirms a slice of strings for the given CharMap.
func GetCharMap(c CharMap) []string {
	return lookup[c]
}
