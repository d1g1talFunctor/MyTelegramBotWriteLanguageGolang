package main

import (
	"fmt"
	_ "fmt"
	tu "github.com/mymmrac/telego/telegoutil"
	"os"
	_ "os"

	"github.com/mymmrac/telego"
)

func main() {
	botToken := "6630158050:AAGEYPrD0ZnI7prFarQUQ3bf_pmCww5tKmg"
	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	updates, _ := bot.UpdatesViaLongPolling(nil)

	defer bot.StopLongPolling()

	for update := range updates {
		if update.Message != nil {
			chatID := tu.ID(update.Message.Chat.ID)

			keyboard := tu.Keyboard(
				tu.KeyboardRow(tu.KeyboardButton("Старт"),
					tu.KeyboardButton("Помощь"), tu.KeyboardButton("Тех. поддержка")),

				tu.KeyboardRow(tu.KeyboardButton("Отправить локацию").WithRequestLocation(),
					tu.KeyboardButton("Отправить контакт").WithRequestContact(), tu.KeyboardButton("Отмена"),
				),
			)

			message := tu.Message(
				chatID, "С этим сообшением придет клавиатуре").WithReplyMarkup(keyboard)

			_, _ = bot.SendMessage(message)
		}
	}
}
