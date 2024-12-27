package water

import "fyne.io/fyne/v2"

// Maybe start a new goroutine, which is communicating with the temp poller through a channel

func Notify(temp float64, app fyne.App) {
	if temp > 45 {
		var notif = fyne.NewNotification("Water tank", "Water tank temperature has exceeded 45 degrees. Maybe turn the heater off?")
		app.SendNotification(notif)
	}
}
