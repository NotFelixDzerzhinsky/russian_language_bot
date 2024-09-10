package telegram

import (
	"RusLangTgBot/database"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

type Bot struct {
	bot    *tgbotapi.BotAPI
	states map[int]*StateMachine
	db     *database.Database
}

func NewBot(bot *tgbotapi.BotAPI, db *database.Database) *Bot {
	return &Bot{bot: bot, states: make(map[int]*StateMachine, 100000), db: db}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(u)
	if err != nil {
		return err
	}
	for update := range updates {
		if update.Message == nil {
			continue
		}

		err = b.handler(update.Message)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *Bot) handler(message *tgbotapi.Message) error {
	if _, ok := b.states[message.From.ID]; !ok {
		b.states[message.From.ID] = NewStateMachine()
	}
	if err := b.db.CheckUserExists("leaderboard",
		message.From.ID, message.From.UserName); err != nil {
		return err
	}
	for _, task := range TasksNumbers {
		if err := b.db.CheckUserExists(fmt.Sprintf("leaderboard%v", task),
			message.From.ID, message.From.UserName); err != nil {
			return err
		}
	}

	if message.IsCommand() {
		if err := b.handleCommand(message); err != nil {
			return err
		}
		return nil
	}

	if err := b.handleMessage(message); err != nil {
		return err
	}
	return nil
}
