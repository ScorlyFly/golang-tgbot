package run

import (
	"bot/internal/bot/config"
	"bot/internal/bot/notification"
	"log"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func StartWebhook(conf *config.Config, bot *gotgbot.Bot, updater *ext.Updater) {
	webhookDomain := conf.WebhookDomain
	webhookSecret := conf.WebhookSecret

	webhookOpts := ext.WebhookOpts{
		ListenAddr:  conf.WebhookAdress, // This example assumes you're in a dev environment running ngrok on 8080.
		SecretToken: webhookSecret,      // Setting a webhook secret here allows you to ensure the webhook is set by you (must be set here AND in SetWebhook!).
	}

	err := updater.StartWebhook(bot, conf.BotToken, webhookOpts)
	if err != nil {
		panic("failed to start webhook: " + err.Error())
	}

	err = updater.SetAllBotWebhooks(webhookDomain, &gotgbot.SetWebhookOpts{
		MaxConnections:     100,
		DropPendingUpdates: true,
		SecretToken:        webhookOpts.SecretToken,
	})
	if err != nil {
		panic("failed to set webhook: " + err.Error())
	}

	log.Printf("%s has been started...\n", bot.User.Username)

	notification.NotifStart(bot)

	// Idle, to keep updates coming in, and avoid bot stopping.
	updater.Idle()
}
