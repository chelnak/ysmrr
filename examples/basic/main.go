package main

import (
	"time"

	"github.com/chelnak/ysmrr"
)

func main() {
	// Create a new spinner manager
	sm := ysmrr.NewSpinnerManager()

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
