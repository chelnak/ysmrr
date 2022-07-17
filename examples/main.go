package main

import (
	"time"

	"github.com/chelnak/ysmrr"
)

func main() {
	// Create a new spinner manager
	sm := ysmrr.NewSpinnerManager()

	// Set up our spinners
	downloading := sm.AddSpinner("Downloading...")
	installing := sm.AddSpinner("Installing...")
	running := sm.AddSpinner("Running...")

	// Start the spinners that have been added to the group
	sm.Start()

	// Set downloading to complete
	time.Sleep(2 * time.Second)
	downloading.Complete()

	// Update the message of the installing spinner
	time.Sleep(2 * time.Second)
	installing.Update("Installing updated...")

	// Set installing to complete
	time.Sleep(2 * time.Second)
	installing.Complete()

	// Set running to error
	time.Sleep(2 * time.Second)
	running.Error()

	// Stop the spinners in the group
	time.Sleep(2 * time.Second)
	sm.Stop()
}
