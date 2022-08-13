// Package animations provides a collection of spinner animations.
// Animations have been borrowed from the following sources:
// * https://wiki.tcl-lang.org/page/Text+Spinner
// * https://stackoverflow.com/questions/2685435/cooler-ascii-spinners
package animations

type Animation int

const (
	Arc Animation = iota
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

var lookup = map[Animation][]string{
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

// GetAnimation retirms a slice of strings for the given type.
func GetAnimation(a Animation) []string {
	return lookup[a]
}
