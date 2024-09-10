package telegram

import (
	"RusLangTgBot/ranks"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) defaultCommand(message *tgbotapi.Message) (string, error) {
	currentState := b.states[message.From.ID].Current()
	if currentState == stateDefault {
		return startMessage, nil
	} else {
		return unwantedCommandMessage, nil
	}
}

func (b *Bot) taskCommand(message *tgbotapi.Message) (string, error) {
	currentState := b.states[message.From.ID].Current()
	if currentState == stateDefault {
		err := b.states[message.From.ID].Event(eventCommandTask)
		if err != nil {
			return "", err
		}
		return getNumberMessage, nil
	} else {
		return unwantedCommandMessage, nil
	}
}

func (b *Bot) trainCommand(message *tgbotapi.Message) (string, error) {
	currentState := b.states[message.From.ID]
	if currentState.Current() == stateDefault {
		err := currentState.Event(eventCommandTrain)
		if err != nil {
			return "", err
		}
		return getNumberMessage, nil
	} else {
		return unwantedCommandMessage, nil
	}
}

func (b *Bot) leaderboardCommand(message *tgbotapi.Message) (string, error) {
	currentState := b.states[message.From.ID]
	if currentState.Current() == stateDefault {
		err := currentState.Event(eventCommandLeaderboard)
		if err != nil {
			return "", err
		}
		return getNumberMessage, nil
	} else {
		return unwantedCommandMessage, nil
	}
}

func (b *Bot) exitCommand(message *tgbotapi.Message) (string, error) {
	currentState := b.states[message.From.ID]
	if currentState.Current() == stateDefault {
		return unwantedTextMessage, nil
	} else {
		err := currentState.Event(eventCommandExit)
		if err != nil {
			return "", err
		}
		return exitCommandMessage, nil
	}
}

func (b *Bot) ranksCommand(message *tgbotapi.Message) (string, error) {
	return ranks.GetRanksTable(), nil
}
