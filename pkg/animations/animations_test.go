package animations_test

import (
	"testing"

	"github.com/chelnak/ysmrr/pkg/animations"
	"github.com/stretchr/testify/assert"
)

var (
	Arc            = []string{"◜", "◠", "◝", "◞", "◡", "◟"}
	Arrow          = []string{"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"}
	Baloon         = []string{".", "o", "O", "@", "*"}
	Baloon2        = []string{".", "o", "O", "°", "O", "o", "."}
	Circle         = []string{"◡", "⊙", "◠"}
	CircleHalves   = []string{"◐", "◓", "◑", "◒"}
	CircleQuarters = []string{"◴", "◷", "◶", "◵"}
	Dots           = []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	Hamburger      = []string{"☱", "☲", "☴"}
	Layer          = []string{"-", "=", "≡"}
	Pipe           = []string{"┤", "┘", "┴", "└", "├", "┌", "┬", "┐"}
	Point          = []string{"∙∙∙", "●∙∙", "∙●∙", "∙∙●", "∙∙∙"}
	Star           = []string{"✶", "✸", "✹", "✺", "✹", "✷"}
	SquareCorners  = []string{"◰", "◳", "◲", "◱"}
)

func TestAnimations(t *testing.T) {
	tests := []struct {
		name string
		s    animations.Animation
		want []string
	}{
		{name: "Arc", s: animations.Arc, want: Arc},
		{name: "Arrow", s: animations.Arrow, want: Arrow},
		{name: "Baloon", s: animations.Baloon, want: Baloon},
		{name: "Baloon2", s: animations.Baloon2, want: Baloon2},
		{name: "Circle", s: animations.Circle, want: Circle},
		{name: "CircleHalves", s: animations.CircleHalves, want: CircleHalves},
		{name: "CircleQuarters", s: animations.CircleQuarters, want: CircleQuarters},
		{name: "Dots", s: animations.Dots, want: Dots},
		{name: "Hamburger", s: animations.Hamburger, want: Hamburger},
		{name: "Layer", s: animations.Layer, want: Layer},
		{name: "Pipe", s: animations.Pipe, want: Pipe},
		{name: "Point", s: animations.Point, want: Point},
		{name: "Star", s: animations.Star, want: Star},
		{name: "SquareCorners", s: animations.SquareCorners, want: SquareCorners},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := animations.GetAnimation(tt.s)
			assert.Equal(t, tt.want, got)
		})
	}
}
