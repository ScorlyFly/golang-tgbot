package handlers

import (
	//moder_commands "bot/internal/bot/handlers/moder/commands"
	user_command "bot/internal/bot/handlers/user/commands"
	// user_message "bot/internal/bot/handlers/user/messages"

	admin_command "bot/internal/bot/handlers/admin/commands"

	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
	//"github.com/PaulSonOfLars/gotgbot/v2/ext/handlers/filters"

	admin "bot/internal/bot/filters/admin"
)

func RegisterHandlers(dp *ext.Dispatcher) {

	//user  ---------------------------------------------------------
	//dp.AddHandler(handlers.NewMessage(message.Text, messages.Echo))
	dp.AddHandler(handlers.NewCommand("start", user_command.Start))

	//dp.AddHandler(handlers.NewMessage(moderFilter.FilterCommand("start"), user_message.All))
	// end user -------------------------------------------------------------------

	// Moder  ---------------------------------------------------------
	//dp.AddHandler(handlers.NewMessage(message.Text, messages.Echo))
	//dp.AddHandler(handlers.NewCommand(filters.Message("moder"), moder_commands.Moder))
	// end Moder -------------------------------------------------------------------

	// Admin  ---------------------------------------------------------
	dp.AddHandler(handlers.NewMessage(admin.IsAdmin, admin_command.Admin))
	// End Admin  -----------------------------------------------------
}
