package colors_test

import (
	"testing"

	"github.com/chelnak/ysmrr/pkg/colors"
	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
)

func TestGetColor(t *testing.T) {
	tests := []struct {
		name string
		c    colors.Color
		want *color.Color
	}{
		{
			name: "GetColor returns the correct mapping for FgHiGreen",
			c:    colors.FgHiGreen,
			want: color.New(color.FgHiGreen),
		},
		{
			name: "GetColor returns the correct mapping for FgHiYellow",
			c:    colors.FgHiYellow,
			want: color.New(color.FgHiYellow),
		},
		{
			name: "GetColor returns the corect mapping for FgHiBlue",
			c:    colors.FgHiBlue,
			want: color.New(color.FgHiBlue),
		},
		{
			name: "GetColor returns the correct mapping for FgHiRed",
			c:    colors.FgHiRed,
			want: color.New(color.FgHiRed),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, colors.GetColor(tt.c))
		})
	}
}
