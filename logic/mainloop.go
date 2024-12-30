package logic

import (
	"fyne.io/fyne/v2"
	"time"
	"watertemp/config"
	"watertemp/water"
)

func (app *App) MainLoop() {
	var hasFailed = false
	var userNotifiedOnce = false

	// This will continuously set the temperature label every second with updated offset information
	go func() {
		for {
			app.tankStatus.LastSuccessfulPollSeconds = app.calculateSuccessfulPollOffset()
			app.temperatureLabel.Label = app.tankStatus.String()
			app.trayMenu.Refresh()

			time.Sleep(time.Second)
		}
	}()

	// Code for fetching the temperature and handling failure
	for {
		var temp, err = water.GetWaterTemp(*config.ResourceIP)
		if err != nil && !hasFailed {
			var recentFailLabel = fyne.NewMenuItem("Failed to poll water tank temperature. Last know value shown", nil)
			recentFailLabel.Disabled = true

			// This is a quirk with Fyne. I cannot append() this failure label to the system tray as that creates another extra "Quit" button
			var itemReconstruction = []*fyne.MenuItem{
				app.temperatureLabel,
				recentFailLabel,
			}
			app.trayMenu.Items = itemReconstruction

			app.trayMenu.Refresh()

			hasFailed = true
		}

		// When temperature is polled correctly
		if err == nil {
			app.tankStatus.Temperature = temp
			hasFailed = false
			app.successfulPollTimestamp = time.Now().Unix()

			app.trayMenu.Items = []*fyne.MenuItem{
				app.temperatureLabel,
			}
			app.trayMenu.Refresh()

			// Notification
			if !userNotifiedOnce {
				water.Notify(app.tankStatus.Temperature, app.fyneApp)
				userNotifiedOnce = true
			} else if app.tankStatus.Temperature < 45 {
				userNotifiedOnce = false // so we can send another notification when it exceeds 45 eventually
			}

			if *config.Debug {
				water.Notify(60, app.fyneApp)
			}
		}

		if *config.Debug {
			time.Sleep(time.Second * 10)
		} else {
			time.Sleep(time.Minute)
		}
	}
}

func (app *App) calculateSuccessfulPollOffset() int64 {
	var nowTime = time.Now().Unix()

	if app.successfulPollTimestamp == 0 {
		return 0
	} else {
		return nowTime - app.successfulPollTimestamp
	}
}
