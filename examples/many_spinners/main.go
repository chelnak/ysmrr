package main

import (
	"strconv"
	"sync"
	"time"

	"github.com/chelnak/ysmrr"
)

func main() {
	// Create a new spinner manager
	sm := ysmrr.NewSpinnerManager()
	sm.Start()
	defer sm.Stop()

	waitGroup := sync.WaitGroup{}
	for i := 0; i < 50; i++ {
		waitGroup.Add(1)
		go func(i int) {
			defer waitGroup.Done()
			s := sm.AddSpinner("Downloading... " + strconv.Itoa(i))
			time.Sleep(20 * time.Second)
			s.Complete()
		}(i)
	}
	waitGroup.Wait()
}
