package app

import (
	"runtime"

	"github.com/islu/crypto-bar/internal/usecase"
	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/foundation"
)

func Launch() {
	runtime.LockOSThread()
	application := NewApplication()
	application.AA.Run()
}

type Application struct {
	AA            appkit.Application
	TickerManager usecase.TickerManager
}

func NewApplication() *Application {

	app := appkit.Application_SharedApplication()
	ticker := usecase.NewTickerManager()

	delegate := &appkit.ApplicationDelegate{}
	delegate.SetApplicationDidFinishLaunching(func(foundation.Notification) {
		app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
		app.ActivateIgnoringOtherApps(true)

	})
	delegate.SetApplicationWillFinishLaunching(func(foundation.Notification) {
		// Set up the system bar
		setMenuBar(app, &ticker)
	})
	delegate.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
		return false
	})

	app.SetDelegate(delegate)

	return &Application{
		AA:            app,
		TickerManager: ticker,
	}
}
