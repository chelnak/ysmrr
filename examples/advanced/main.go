package main

import (
	"sync"
	"time"

	"github.com/chelnak/ysmrr"
	"github.com/chelnak/ysmrr/pkg/animations"
	"github.com/chelnak/ysmrr/pkg/colors"
)

func main() {
	// Create a new spinner manager
	sm := ysmrr.NewSpinnerManager(
		ysmrr.WithAnimation(animations.Pipe),
		ysmrr.WithSpinnerColor(colors.FgHiBlue),
	)

	// Set up our spinners
	downloading := sm.AddSpinner("Downloading...")
	installing := sm.AddSpinner("Installing...")
	running := sm.AddSpinner("Running...")

	// Start the spinners that have been added to the group
	sm.Start()
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

func downloader(wg *sync.WaitGroup, s *ysmrr.Spinner) {
	files := []string{"file1", "file2", "file3", "file4", "file5"}
	for _, file := range files {
		s.UpdateMessagef("Downloading %s...", file)
		time.Sleep(1 * time.Second)
	}

	s.UpdateMessage("Downloading complete...")
	s.Complete()
}

func installer(wg *sync.WaitGroup, s *ysmrr.Spinner) {
	s.UpdateMessage("Installing...")
	time.Sleep(4 * time.Second)
	s.UpdateMessage("Installation complete...")
	s.Complete()
}

func runner(wg *sync.WaitGroup, s *ysmrr.Spinner) {
	s.UpdateMessage("Running...")
	time.Sleep(8 * time.Second)
	s.Error()
	s.UpdateMessage("Running encountered an error...")
}
