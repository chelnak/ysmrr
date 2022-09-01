package animations_test

import (
	"testing"
	"time"

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
		name      string
		s         animations.Animation
		wantChars []string
		wantSpeed time.Duration
	}{
		{name: "Arc", s: animations.Arc, wantChars: Arc, wantSpeed: 100},
		{name: "Arrow", s: animations.Arrow, wantChars: Arrow, wantSpeed: 100},
		{name: "Baloon", s: animations.Baloon, wantChars: Baloon, wantSpeed: 140},
		{name: "Baloon2", s: animations.Baloon2, wantChars: Baloon2, wantSpeed: 120},
		{name: "Circle", s: animations.Circle, wantChars: Circle, wantSpeed: 120},
		{name: "CircleHalves", s: animations.CircleHalves, wantChars: CircleHalves, wantSpeed: 50},
		{name: "CircleQuarters", s: animations.CircleQuarters, wantChars: CircleQuarters, wantSpeed: 120},
		{name: "Dots", s: animations.Dots, wantChars: Dots, wantSpeed: 80},
		{name: "Hamburger", s: animations.Hamburger, wantChars: Hamburger, wantSpeed: 100},
		{name: "Layer", s: animations.Layer, wantChars: Layer, wantSpeed: 150},
		{name: "Pipe", s: animations.Pipe, wantChars: Pipe, wantSpeed: 100},
		{name: "Point", s: animations.Point, wantChars: Point, wantSpeed: 125},
		{name: "Star", s: animations.Star, wantChars: Star, wantSpeed: 70},
		{name: "SquareCorners", s: animations.SquareCorners, wantChars: SquareCorners, wantSpeed: 180},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSpeed, gotChars := animations.GetAnimation(tt.s)
			assert.Equal(t, tt.wantChars, gotChars)
			assert.Equal(t, tt.wantSpeed*time.Millisecond, gotSpeed)
		})
	}
}

var properties = animations.Properties{
	Speed:      100,
	Characters: []string{"a", "b", "c"},
}

func TestProperties_GetSpeed(t *testing.T) {
	assert.Equal(t, 100*time.Millisecond, properties.GetSpeed())
}

func TestProperties_GetCharacters(t *testing.T) {
	assert.Equal(t, []string{"a", "b", "c"}, properties.GetCharacters())
}

func TestGetAnimations(t *testing.T) {
	got := animations.GetAnimations()
	assert.IsType(t, []animations.Animation{}, got)
	assert.Equal(t, 14, len(got))
}
