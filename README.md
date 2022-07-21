# You spin me right round

Ysmrr is a simple multi-line spinner package for your terminal.

!["ysmrr - examples/advanced/main.go"](advanced-example.gif)

## Installing

```bash
go get -u github.com/chelnak/ysmrr
```

## Usage

``` go
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

For example usage, check out the [examples](examples) directory for.

## Inspiration

Ysmrr was inspired by the following projects:

* [github.com/briandowns/spinner](https://github.com/briandowns/spinner)
* [github.com/theckman/yacspin](https://github.com/theckman/yacspin)

It also uses [github.com/fatih/color](https://github.com/fatih/color) for the underlying color system
and [github.com/mattn/go-colorable](https://github.com/mattn/go-colorable) for Windows support.
