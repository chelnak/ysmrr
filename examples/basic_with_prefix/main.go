package main

import (
	"time"

	"github.com/chelnak/ysmrr"
	"github.com/chelnak/ysmrr/pkg/animations"
	"github.com/chelnak/ysmrr/pkg/colors"
)

func main() {
	// Create a new spinner manager
	sm := ysmrr.NewSpinnerManager(
		ysmrr.WithAnimation(animations.Arrow),
		ysmrr.WithSpinnerColor(colors.FgHiBlue),
		ysmrr.WithMessageColor(colors.FgHiYellow),
	)

	// Set up our spinner
	example := sm.AddSpinner("this text is after the spinner")

	// Add a prefix
	example.UpdatePrefix("this text is before the spinner")

	// Start the spinners that have been added to the group
	sm.Start()
	defer sm.Stop()

	// Set spinner to complete
	time.Sleep(2 * time.Second)
	example.Complete()
}
