package charmap_test

import (
	"testing"

	"github.com/chelnak/ysmrr/pkg/charmap"
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

func TestCharMaps(t *testing.T) {
	tests := []struct {
		name string
		c    charmap.CharMap
		want []string
	}{
		{name: "Arc", c: charmap.Arc, want: Arc},
		{name: "Arrow", c: charmap.Arrow, want: Arrow},
		{name: "Baloon", c: charmap.Baloon, want: Baloon},
		{name: "Baloon2", c: charmap.Baloon2, want: Baloon2},
		{name: "Circle", c: charmap.Circle, want: Circle},
		{name: "CircleHalves", c: charmap.CircleHalves, want: CircleHalves},
		{name: "CircleQuarters", c: charmap.CircleQuarters, want: CircleQuarters},
		{name: "Dots", c: charmap.Dots, want: Dots},
		{name: "Hamburger", c: charmap.Hamburger, want: Hamburger},
		{name: "Layer", c: charmap.Layer, want: Layer},
		{name: "Pipe", c: charmap.Pipe, want: Pipe},
		{name: "Point", c: charmap.Point, want: Point},
		{name: "Star", c: charmap.Star, want: Star},
		{name: "SquareCorners", c: charmap.SquareCorners, want: SquareCorners},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := charmap.GetCharMap(tt.c)
			assert.Equal(t, tt.want, got)
		})
	}
}
