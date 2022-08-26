// Package tput provides convenience functions for sending escape sequences to the terminal.
// The escape codes used have been derrived from the tput program.
package tput

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mattn/go-isatty"
)

// Sc saves the current position of the cursor.
func Sc(w io.Writer) {
	write(w, "\u001b7")
}

// Rc restores the cursor to the saved position.
func Rc(w io.Writer) {
	write(w, "\u001b8")
}

// Civis hides the cursor.
func Civis(w io.Writer) {
	write(w, "\u001b[?25l")
}

// Cnorm shows the cursor.
func Cnorm(w io.Writer) {
	writef(w, "\u001b[?25h")
}

// Cuu moves the cursor up by n lines.
func Cuu(w io.Writer, n int) {
	writef(w, "\u001b[%dA", n)
}

// BufScreen ensures that there are enough lines available
// by sending n * newlines to the writer.
func BufScreen(w io.Writer, n int) {
	writef(w, "%s", strings.Repeat("\n", n))
}

func ClearLine(w io.Writer) {
	write(w, "\u001b[K")
}

func tty() bool {
	return isatty.IsTerminal(os.Stdout.Fd()) || os.Getenv("YSMRR_FORCE_TTY") == "true"
}

func write(w io.Writer, s string) {
	if tty() {
		fmt.Fprint(w, s)
	}
}

func writef(w io.Writer, format string, a ...interface{}) {
	if tty() {
		fmt.Fprintf(w, format, a...)
	}
}
