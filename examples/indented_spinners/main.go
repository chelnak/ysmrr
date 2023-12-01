package main

import (
	"strconv"
	"time"

	"github.com/chelnak/ysmrr"
)

func main() {
	// Create a new spinner manager
	sm := ysmrr.NewSpinnerManager()

	// Setup a group of spinners
	spinners := make(map[int]*ysmrr.Spinner, 5)
	spinners[0] = sm.AddSpinner("Worker Group")
	for i := 1; i <= 4; i++ {
		spinners[i] = sm.AddSpinner("worker" + strconv.Itoa(i))

		// Indent worker spinners
		spinners[i].UpdatePrefix("  ")
	}

	// Start the spinners that have been added to the group
	sm.Start()
	defer sm.Stop()

	// Simulate work
	for i := 1; i <= 4; i++ {
		time.Sleep(2 * time.Second)
		spinners[i].Complete()
	}

	// All the workers are done, set status of group
	spinners[0].Complete()
}
