// Package tput provides convenience functions for sending escape sequences to the terminal.
// The escape codes used have been derrived from the tput program.
package tput

import "fmt"

// Sc saves the current position of the cursor
func Sc() {
	fmt.Printf("\u001b7")
}

// Rc restores the cursor to the saved position
func Rc() {
	fmt.Printf("\u001b8")
}

// Civis hides the cursor
func Civis() {
	fmt.Printf("\u001b[?25l")
}

// Cnorm shows the cursor
func Cnorm() {
	fmt.Printf("\u001b[?12l[?25h")
}
