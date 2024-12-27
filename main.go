package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"log"
	"time"
)

const debug = false
const localIp = "10.50.0.116"

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
	fyneApp.SetIcon(theme.FyneLogo())

	// Type assertion https://go.dev/tour/methods/15
	if trayControl, isDesktop = fyneApp.(desktop.App); !isDesktop {
		log.Fatal("The environment in which this app is running isn't a desktop. This is a desktop application as it requires the system tray")
	}

	trayControl.SetSystemTrayMenu(temperatureMenu)

	go mainLoop()
	fyneApp.Run()
	// fyneApp.Run is blocking
}

func mainLoop() {
	var hasFailed = false
	var tankTemp float64 = 0

	// This will continuously set the temperature label every second with updated offset information
	go func() {
		for {
			temperatureLabel.Label = fmt.Sprintf("Water tank temperature is %.2f (%d seconds ago)", tankTemp, calculateSuccessfulPollOffset())
			temperatureMenu.Refresh()

			time.Sleep(time.Second)
		}
	}()

	// Code for fetching the temperature and handling failure
	for {
		var temp, err = getWaterTemp()
		if err != nil && !hasFailed {
			var recentFailLabel = fyne.NewMenuItem("Failed to poll water tank temperature. Last know value shown", nil)
			recentFailLabel.Disabled = true

			// This is a quirk with Fyne. I cannot append this failure label to the system tray as that creates another extra "Quit" button
			var itemReconstruction = []*fyne.MenuItem{
				temperatureLabel,
				recentFailLabel,
			}
			temperatureMenu.Items = itemReconstruction

			temperatureMenu.Refresh()

			hasFailed = true
		}

		// When temperature is polled correctly
		if err == nil {
			tankTemp = temp
			hasFailed = false
			successfulPollTimestamp = time.Now().Unix()

			temperatureMenu.Items = []*fyne.MenuItem{
				temperatureLabel,
			}
			temperatureMenu.Refresh()
		}

		if debug {
			time.Sleep(time.Second * 10)
		} else {
			time.Sleep(time.Minute)
		}
	}
}

func calculateSuccessfulPollOffset() int64 {
	var nowTime = time.Now().Unix()

	if successfulPollTimestamp == 0 {
		return 0
	} else {
		return nowTime - successfulPollTimestamp
	}
}
