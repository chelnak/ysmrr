// Package tput provides convenience functions for sending escape sequences to the terminal.
// The escape codes used have been derrived from the tput program.
package tput

import (
	"fmt"
	"io"
)

// Sc saves the current position of the cursor.
func Sc(w io.Writer) {
	fmt.Fprint(w, "\u001b7")
}

// Rc restores the cursor to the saved position.
func Rc(w io.Writer) {
	fmt.Fprint(w, "\u001b8")
}

// Civis hides the cursor.
func Civis(w io.Writer) {
	fmt.Fprint(w, "\u001b[?25l")
}

// Cnorm shows the cursor.
func Cnorm(w io.Writer) {
	fmt.Fprintf(w, "\u001b[?25h")
}
