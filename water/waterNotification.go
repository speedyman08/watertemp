package water

import (
	"fyne.io/fyne/v2"
	"math"
)

func Notify(temp float64, app fyne.App) {
	if math.Round(temp) > 45 {
		var notif = fyne.NewNotification("Water tank", "Water tank temperature has exceeded 45 degrees. Maybe turn the heater off?")
		app.SendNotification(notif)
	}
}
