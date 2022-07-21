package tput_test

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/chelnak/ysmrr/pkg/tput"
	"github.com/stretchr/testify/assert"
)

func TestTput(t *testing.T) {
	tests := []struct {
		name string
		fn   func(w io.Writer)
		want string
	}{
		{
			name: "Sc",
			fn:   tput.Sc,
			want: "\u001b7",
		},
		{
			name: "Rc",
			fn:   tput.Rc,
			want: "\u001b8",
		},
		{
			name: "Civis",
			fn:   tput.Civis,
			want: "\u001b[?25l",
		},
		{
			name: "Cnorm",
			fn:   tput.Cnorm,
			want: "\u001b[?25h",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			tt.fn(&buf)
			assert.Equal(t, tt.want, buf.String())
		})
	}
}

func TestTputCommandsWithInputs(t *testing.T) {
	tests := []struct {
		name  string
		input int
		fn    func(w io.Writer, n int)
		want  string
	}{
		{
			name:  "Cuu",
			input: 1,
			fn:    tput.Cuu,
			want:  "\u001b[1A",
		},
		{
			name:  "BufScreen",
			input: 10,
			fn:    tput.BufScreen,
			want:  strings.Repeat("\n", 10),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			tt.fn(&buf, tt.input)
			assert.Equal(t, tt.want, buf.String())
		})
	}
}
