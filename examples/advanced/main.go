package main

import (
	"fmt"
	"sync"
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
	sm.Init()
	defer sm.Stop()

	// Set up our wait group
	var wg sync.WaitGroup
	wg.Add(len(sm.GetSpinners()))

	go func() {
		defer wg.Done()
		downloader(&wg, downloading)
	}()

	go func() {
		defer wg.Done()
		installer(&wg, installing)
	}()

	go func() {
		defer wg.Done()
		runner(&wg, running)
	}()

	// Wait for all of the workders to complete
	wg.Wait()
}

// Simulate some work
func calculateFibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return calculateFibonacci(n-1) + calculateFibonacci(n-2)
}

func downloader(wg *sync.WaitGroup, s *ysmrr.Spinner) {
	time.Sleep(2 * time.Second)

	s.Start()
	for i := 0; i < 10; i++ {
		n := calculateFibonacci(i)
		message := fmt.Sprintf("Downloading %d...", n)
		s.UpdateMessage(message)
		time.Sleep(1 * time.Second)
	}

	s.UpdateMessage("Download complete...")
	s.Complete()
}

func installer(wg *sync.WaitGroup, s *ysmrr.Spinner) {
	time.Sleep(5 * time.Second)
	s.Start()

	for i := 0; i < 10; i++ {
		n := calculateFibonacci(i)
		message := fmt.Sprintf("Installing %d...", n)
		s.UpdateMessage(message)
		time.Sleep(200 * time.Millisecond)
	}

	s.UpdateMessage("Installation complete...")
	s.Complete()
}

func runner(wg *sync.WaitGroup, s *ysmrr.Spinner) {
	s.Start()
	s.UpdateMessage("Running...")
	time.Sleep(8 * time.Second)
	s.Error()
	s.UpdateMessage("Running encountered an error...")
}
