package main

import (
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"log"
)

const debug = false
const localIp = "10.50.0.116"

// thermometer icon for the system tray
//
//go:embed resources/thermometer.png
var iconBytes []byte

var successfulPollTimestamp int64 = 0

var temperatureLabel = fyne.NewMenuItem("Placeholder (Wait)", nil)
var temperatureMenu = fyne.NewMenu("Temperature App", temperatureLabel)

func main() {
	temperatureLabel.Disabled = true

	var (
		fyneApp     = app.New()
		trayControl desktop.App
		isDesktop   bool
	)

	// Setting icon
	var rss = fyne.NewStaticResource("thermometer.png", iconBytes)

	fyneApp.SetIcon(rss)

	// Type assertion https://go.dev/tour/methods/15
	if trayControl, isDesktop = fyneApp.(desktop.App); !isDesktop {
		log.Fatal("The environment in which this app is running isn't a desktop. This is a desktop application as it requires the system tray")
	}

	trayControl.SetSystemTrayMenu(temperatureMenu)

	go mainLoop()
	fyneApp.Run()
	// fyneApp.Run is blocking
}
