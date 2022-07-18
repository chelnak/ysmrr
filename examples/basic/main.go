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
	sm.Start()
	defer sm.Stop()

	// Set downloading to complete
	time.Sleep(2 * time.Second)
	downloading.Complete()
}
