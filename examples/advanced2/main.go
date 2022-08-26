package main

import (
	"time"

	"github.com/chelnak/ysmrr"
)

type App struct {
	spinnerManager ysmrr.SpinnerManager
}

func (a *App) doMoreWork() {
	s := a.spinnerManager.AddSpinner("Loading...DoMoreWork()")
	time.Sleep(time.Second * 5)
	s.UpdateMessage("Done")
	s.Complete()
}

func (a *App) doWork() {

	s := a.spinnerManager.AddSpinner("Loading...DoWork()")
	time.Sleep(time.Second * 5)
	s.UpdateMessage("its Done")
	s.Complete()

	a.doMoreWork()
}

func (a *App) Run() {
	a.spinnerManager.Start()
	s := a.spinnerManager.AddSpinner("Loading...Run()")
	a.doWork()
	s.UpdateMessage("Loading...Run() Done")
	s.Complete()
}

func (a *App) Stop() {
	defer a.spinnerManager.Stop()
}

func NewApp() *App {
	return &App{
		spinnerManager: ysmrr.NewSpinnerManager(),
	}
}

func main() {

	app := NewApp()
	app.Run()
	app.Stop()

}
