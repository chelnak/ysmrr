package main

import (
	"time"

	"github.com/chelnak/ysmrr"
	"github.com/chelnak/ysmrr/pkg/charmap"
	"github.com/chelnak/ysmrr/pkg/colors"
)

func main() {
	// Create a new spinner manager
	sm := ysmrr.NewSpinnerManager(
		ysmrr.WithCharMap(charmap.Arrows),
		ysmrr.WithSpinnerColor(colors.FgHiBlue),
		ysmrr.WithMessageColor(colors.FgHiYellow),
	)

	// Set up our spinner
	downloading := sm.AddSpinner("Downloading...")

	// Start the spinners that have been added to the group
	sm.Init()
	defer sm.Stop()

	downloading.Start()
	time.Sleep(2 * time.Second)

	// Set downloading to complete
	downloading.Complete()
}
