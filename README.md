# You spin me right round

YSMRR is a simple multi-line spinner package for your terminal.

## Installing

```bash
go get -u github.com/chelnak/ysmrr
```

## Usage


```bash
  // Create a new spinner manager
  sm := ysmrr.NewSpinnerManager()

	// Add a spinner
	mySpinner := sm.AddSpinner("Spinny things...")

	// Start the spinners that have been added to the group
	sm.Start()

	// Set the spinner to complete
	time.Sleep(2 * time.Second)
	mySpinner.Complete()

  // Stop the spinners in the group
	time.Sleep(2 * time.Second)
	sm.Stop()
```

See [examples/main.go](examples/main.go) for a more in-depth example.
