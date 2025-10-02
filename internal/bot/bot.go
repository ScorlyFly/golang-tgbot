package bot

import (
	"bot/internal/bot/config"
	botHandler "bot/internal/bot/handlers"
	"bot/internal/bot/notification"
	"bot/internal/bot/run"
	"log"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func Run(getConfig string) {
	// загружаем конфиг
	config.Conf = config.LoadConfig(getConfig)
	// создаем бота
	bot := initBot(config.Conf.BotToken)
	// инициализирует Dispatcher
	dp := Dispatcher()
	// на основе диспетчера разбиваем на роли (routers)
	//создадим обработчик update
	updater := ext.NewUpdater(dp, nil)
	botHandler.RegisterHandlers(dp)

	//run.StartPolling(bot, updater)
	run.StartWebhook(config.Conf, bot, updater)
	
	notification.NotifStop(bot)
}

func initBot(token string) *gotgbot.Bot {
	// создаём бота
	bot, err := gotgbot.NewBot(token, nil)
	if err != nil {
		log.Fatal("ERROR: bot not created ", err)
	} else {
		log.Printf("INFO: bot success created ", bot.User.Username)
	}
	return bot
}

func Dispatcher() *ext.Dispatcher {
	// Create updater and dispatcher.
	initDp := ext.NewDispatcher(&ext.DispatcherOpts{
		// If an error is returned by a handler, log it and continue going.
		Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
			log.Println("an error occurred while handling update:", err.Error())
			return ext.DispatcherActionNoop
		},
		MaxRoutines: ext.DefaultMaxRoutines,
	})
	return initDp
}
