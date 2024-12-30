package main

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"log"
	"watertemp/water"
)

// thermometer icon for the system tray
//
//go:embed resources/thermometer.png
var iconBytes []byte

var globals = Globals{}

type Globals struct {
	debug      bool
	resourceIp string
	iconBytes  []byte
}

type App struct {
	fyneApp          fyne.App
	trayMenu         *fyne.Menu
	temperatureLabel *fyne.MenuItem
	// The last unix time the tank temperature was polled
	successfulPollTimestamp int64
	tankStatus              water.Status
}

func NewApp() (a *App) {
	a = new(App)
	a.fyneApp = app.New()
	a.temperatureLabel = fyne.NewMenuItem("Placeholder (Wait)", nil)
	a.trayMenu = fyne.NewMenu("Temperature App", a.temperatureLabel)

	a.temperatureLabel.Disabled = true

	return
}

func (app *App) setTray() {
	var (
		trayControl desktop.App
		isDesktop   bool
	)

	// Setting icon
	var rss = fyne.NewStaticResource("thermometer.png", iconBytes)

	app.fyneApp.SetIcon(rss)

	// Type assertion https://go.dev/tour/methods/15
	if trayControl, isDesktop = app.fyneApp.(desktop.App); !isDesktop {
		log.Fatal("The environment in which this app is running isn't a desktop. This is a desktop application as it requires the system tray")
	}

	trayControl.SetSystemTrayMenu(app.trayMenu)
}

func (app *App) Run() {
	app.setTray()
	go app.MainLoop()
	app.fyneApp.Run()
}

func main() {
	globals.debug = false
	globals.resourceIp = "10.50.0.116"
	globals.iconBytes = iconBytes

	var a = NewApp()
	a.Run()
}
