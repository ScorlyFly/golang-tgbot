package admin

import (
	"bot/internal/bot/config"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

func IsAdmin(msg *gotgbot.Message) bool {
	if msg.Text == "/admin" && msg.Chat.Id == config.Conf.AdminID {
		return true
	} else {
		return false
	}
}
