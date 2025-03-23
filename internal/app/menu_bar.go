package app

import (
	"fmt"
	"time"

	"github.com/islu/crypto-bar/internal/usecase"
	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/objc"
)

func setMenuBar(app appkit.Application, tm *usecase.TickerManager) {
	main := appkit.StatusBar_SystemStatusBar().StatusItemWithLength(appkit.VariableStatusItemLength)
	objc.Retain(&main)

	// Initial title
	tickerPrice := tm.GetTickerPrice(tm.Symbols[0])
	main.Button().SetTitle(tickerPrice + " | " + tm.GetCurrSymbol())

	// Setup the menu
	menu := appkit.NewMenuWithTitle("main")

	// Setup the options menu
	menu.AddItem(setUpdateIntervalBar(tm))
	menu.AddItem(appkit.MenuItem_SeparatorItem())
	// menu.AddItem(appkit.NewMenuItemWithAction("Hide", "h", func(sender objc.Object) { app.Hide(nil) }))
	menu.AddItem(appkit.NewMenuItemWithAction("Quit", "q", func(sender objc.Object) { app.Terminate(nil) }))

	main.SetMenu(menu)

	// Update the ticker price
	go func() {
		ticker := time.NewTicker(tm.GetCurrInterval())
		for range ticker.C {
			fmt.Println("Current interval: ", tm.GetCurrInterval())

			tickerPrice := tm.GetTickerPrice(tm.Symbols[0])
			main.Button().SetTitle(tickerPrice + " | " + tm.GetCurrSymbol())
		}
	}()
}

func setUpdateIntervalBar(tm *usecase.TickerManager) appkit.MenuItem {
	optionsMenu := appkit.NewMenu()
	// optionsMenu.AddItem(appkit.NewMenuItemWithAction("Real-Time", "", func(sender objc.Object) { fmt.Println("Real-Time clicked") }))
	optionsMenu.AddItem(appkit.NewMenuItemWithAction("30 seconds", "", func(sender objc.Object) {
		optionsMenu.ItemAtIndex(tm.CurrIntervalIdx).SetState(appkit.ControlStateValueOff)
		tm.CurrIntervalIdx = 0
		optionsMenu.ItemAtIndex(tm.CurrIntervalIdx).SetState(appkit.ControlStateValueOn)
	}))
	optionsMenu.AddItem(appkit.NewMenuItemWithAction("1 minute", "", func(sender objc.Object) {
		optionsMenu.ItemAtIndex(tm.CurrIntervalIdx).SetState(appkit.ControlStateValueOff)
		tm.CurrIntervalIdx = 1
		optionsMenu.ItemAtIndex(tm.CurrIntervalIdx).SetState(appkit.ControlStateValueOn)
	}))
	optionsMenu.AddItem(appkit.NewMenuItemWithAction("5 minutes", "", func(sender objc.Object) {
		optionsMenu.ItemAtIndex(tm.CurrIntervalIdx).SetState(appkit.ControlStateValueOff)
		tm.CurrIntervalIdx = 2
		optionsMenu.ItemAtIndex(tm.CurrIntervalIdx).SetState(appkit.ControlStateValueOn)
	}))
	optionsMenu.AddItem(appkit.NewMenuItemWithAction("15 minutes", "", func(sender objc.Object) {
		optionsMenu.ItemAtIndex(tm.CurrIntervalIdx).SetState(appkit.ControlStateValueOff)
		tm.CurrIntervalIdx = 3
		optionsMenu.ItemAtIndex(tm.CurrIntervalIdx).SetState(appkit.ControlStateValueOn)
	}))
	optionsMenu.AddItem(appkit.NewMenuItemWithAction("30 minutes", "", func(sender objc.Object) {
		optionsMenu.ItemAtIndex(tm.CurrIntervalIdx).SetState(appkit.ControlStateValueOff)
		tm.CurrIntervalIdx = 4
		optionsMenu.ItemAtIndex(tm.CurrIntervalIdx).SetState(appkit.ControlStateValueOn)
	}))
	optionsMenu.ItemAtIndex(tm.CurrIntervalIdx).SetState(appkit.ControlStateValueOn)

	options := appkit.NewMenuItem()
	options.SetTitle("Update Interval")
	options.SetSubmenu(optionsMenu)

	return options
}
