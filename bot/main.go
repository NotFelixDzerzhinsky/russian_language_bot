package main

import (
	"RusLangTgBot/database"
	"RusLangTgBot/tasks"
	"RusLangTgBot/telegram"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var TasksNumbers = []int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22}

func main() {
	bot, err := tgbotapi.NewBotAPI(TelegramBotToken)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	db, err := database.NewDatabase("./database/user_info.db")
	if err != nil {
		log.Panic(err)
	}

	err = db.Init("leaderboard")
	for _, task := range TasksNumbers {
		err = db.Init(fmt.Sprintf("leaderboard%v", task))
		if err != nil {
			log.Panic(err)
		}
	}

	for _, task := range TasksNumbers {
		fmt.Println(task)
		err = tasks.Init(fmt.Sprintf("./tasks/numbers/task%v/task%v.csv", task, task), task)
		if err != nil {
			log.Panic(err)
		}
	}

	telegramBot := telegram.NewBot(bot, db)
	err = telegramBot.Start()
	if err != nil {
		log.Fatal(err)
	}
}
