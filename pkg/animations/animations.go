// Package animations provides a collection of spinner animations.
// Animations have been borrowed from the following sources:
// * https://wiki.tcl-lang.org/page/Text+Spinner
// * https://stackoverflow.com/questions/2685435/cooler-ascii-spinners
package animations

import (
	"time"
)

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

type Properties struct {
	Speed      time.Duration
	Characters []string
}

func (p Properties) GetSpeed() time.Duration {
	return p.Speed * time.Millisecond
}

func (p Properties) GetCharacters() []string {
	return p.Characters
}

var lookup = map[Animation]Properties{
	Arc:            {100, []string{"◜", "◠", "◝", "◞", "◡", "◟"}},
	Arrow:          {100, []string{"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"}},
	Baloon:         {140, []string{".", "o", "O", "@", "*"}},
	Baloon2:        {120, []string{".", "o", "O", "°", "O", "o", "."}},
	Circle:         {120, []string{"◡", "⊙", "◠"}},
	CircleHalves:   {50, []string{"◐", "◓", "◑", "◒"}},
	CircleQuarters: {120, []string{"◴", "◷", "◶", "◵"}},
	Dots:           {80, []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}},
	Hamburger:      {100, []string{"☱", "☲", "☴"}},
	Layer:          {150, []string{"-", "=", "≡"}},
	Pipe:           {100, []string{"┤", "┘", "┴", "└", "├", "┌", "┬", "┐"}},
	Point:          {125, []string{"∙∙∙", "●∙∙", "∙●∙", "∙∙●", "∙∙∙"}},
	Star:           {70, []string{"✶", "✸", "✹", "✺", "✹", "✷"}},
	SquareCorners:  {180, []string{"◰", "◳", "◲", "◱"}},
}

// GetAnimation retirms a slice of strings for the given type.
func GetAnimation(a Animation) (time.Duration, []string) {
	return lookup[a].GetSpeed(), lookup[a].GetCharacters()
}

// GetAnimations returns an unsorted slice of all available animations ids.
func GetAnimations() []Animation {
	keys := make([]Animation, len(lookup))
	i := 0
	for k := range lookup {
		keys[i] = k
		i++
	}
	return keys
}
