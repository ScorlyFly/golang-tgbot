package notification

import (
	"bot/internal/bot/config"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

func NotifStart(b *gotgbot.Bot) {
	_, err := b.SendMessage(config.Conf.AdminID, "bot startup 😉👌", nil)
	if err != nil {
		panic("failed to send message: " + err.Error())
	}
}

func NotifStop(b *gotgbot.Bot) {
	_, err := b.SendMessage(config.Conf.AdminID, "bot stop 😳😰", nil)
	if err != nil {
		panic("failed to send message: " + err.Error())
	}
}
