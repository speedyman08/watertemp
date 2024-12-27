package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"time"
	"watertemp/water"
)

func mainLoop() {
	var hasFailed = false
	var userNotifiedOnce = false
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
		var temp, err = water.GetWaterTemp(localIp)
		if err != nil && !hasFailed {
			var recentFailLabel = fyne.NewMenuItem("Failed to poll water tank temperature. Last know value shown", nil)
			recentFailLabel.Disabled = true

			// This is a quirk with Fyne. I cannot append() this failure label to the system tray as that creates another extra "Quit" button
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

			// Notification
			if !userNotifiedOnce {
				water.Notify(tankTemp, fyneApp)
				userNotifiedOnce = true
			} else if tankTemp < 45 {
				userNotifiedOnce = false // so we can send another notification when it exceeds 45 eventually
			}
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
