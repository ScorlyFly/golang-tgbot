package run

import (
	"bot/internal/bot/notification"
	"log"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

const (
	dropUpdates = true
	limit       = 100
	timeOut     = 9
	httpTimeOut = 10
)

func StartPolling(bot *gotgbot.Bot, updater *ext.Updater) {
	//Start receiving updates.

	err := updater.StartPolling(bot, &ext.PollingOpts{
		DropPendingUpdates: dropUpdates,
		GetUpdatesOpts: &gotgbot.GetUpdatesOpts{
			Timeout: timeOut,
			Limit:   limit,
			RequestOpts: &gotgbot.RequestOpts{
				Timeout: time.Second * httpTimeOut,
			},
		},
	})
	if err != nil {
		panic("failed to start polling: " + err.Error())
	}
	log.Printf("%s has been started...\n", bot.User.Username)

	notification.NotifStart(bot)
	// Idle, to keep updates coming in, and avoid bot stopping.
	updater.Idle()
}
