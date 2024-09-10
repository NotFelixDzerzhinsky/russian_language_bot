package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, message.Command())
	switch message.Command() {
	case commandStart:
		curMsg, err := b.defaultCommand(message)
		if err != nil {
			return err
		}
		msg.Text = curMsg
	case commandHelp:
		curMsg, err := b.defaultCommand(message)
		if err != nil {
			return err
		}
		msg.Text = curMsg
	case commandExit:
		curMsg, err := b.exitCommand(message)
		if err != nil {
			return err
		}
		msg.Text = curMsg
	case commandTask:
		curMsg, err := b.taskCommand(message)
		if err != nil {
			return err
		}
		msg.Text = curMsg
	case commandTrain:
		curMsg, err := b.trainCommand(message)
		if err != nil {
			return err
		}
		msg.Text = curMsg
	case commandLeaderboard:
		curMsg, err := b.leaderboardCommand(message)
		if err != nil {
			return err
		}
		msg.Text = curMsg
	case commandRanks:
		curMsg, err := b.ranksCommand(message)
		if err != nil {
			return err
		}
		msg.Text = curMsg
	}
	// todo если хотим получить число, то надо дропнуть клавиатуру
	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	var messages []string
	var err error

	switch b.states[message.From.ID].Current() {
	case stateDefault:
		messages, err = b.defaultMessage(message)
		if err != nil {
			return err
		}
	case stateWaitTaskNumber:
		messages, err = b.waitTaskNumberMessage(message)
		if err != nil {
			return err
		}
	case stateWaitTrainNumber:
		messages, err = b.waitTrainNumberMessage(message)
		if err != nil {
			return err
		}
	case stateWaitLeaderboardNumber:
		messages, err = b.waitLeaderboardNumberMessage(message)
		if err != nil {
			return err
		}
	case stateWaitTaskAnswer:
		messages, err = b.waitTaskAnswerMessage(message)
		if err != nil {
			return err
		}
	case stateWaitTrainAnswer:
		messages, err = b.waitTrainAnswerMessage(message)
		if err != nil {
			return err
		}
	}

	for _, messageText := range messages {
		msg := tgbotapi.NewMessage(message.Chat.ID, messageText)
		_, err = b.bot.Send(msg)
		if err != nil {
			return err
		}
	}

	return nil
}
