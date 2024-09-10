package telegram

import (
	"RusLangTgBot/ranks"
	"RusLangTgBot/tasks"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strconv"
)

func checkNumber(number string) bool {
	num, err := strconv.Atoi(number)
	if err != nil {
		return false
	}
	for _, task := range TasksNumbers {
		if task == num {
			return true
		}
	}
	return false
}

func (b *Bot) defaultMessage(message *tgbotapi.Message) ([]string, error) {
	return []string{defaultTextMessage}, nil
}

func (b *Bot) waitTaskNumberMessage(message *tgbotapi.Message) ([]string, error) {
	currentState := b.states[message.From.ID]
	if checkNumber(message.Text) {
		err := currentState.Event(eventGetNumber)
		if err != nil {
			return []string{}, err
		}
		currentState.task, err = strconv.Atoi(message.Text)
		if err != nil {
			return []string{}, err
		}
		currentState.number = tasks.GetRandomTask(currentState.task)
		return []string{tasks.GetStatement(currentState.task, currentState.number)}, nil
	} else {
		return []string{wrongNumberMessage}, nil
	}
}

func (b *Bot) waitTrainNumberMessage(message *tgbotapi.Message) ([]string, error) {
	currentState := b.states[message.From.ID]
	if checkNumber(message.Text) {
		err := currentState.Event(eventGetNumber)
		if err != nil {
			return []string{}, err
		}
		currentState.task, err = strconv.Atoi(message.Text)
		if err != nil {
			return []string{}, err
		}
		currentState.number = tasks.GetRandomTask(currentState.task)
		return []string{tasks.GetStatement(currentState.task, currentState.number)}, nil
	} else {
		return []string{wrongNumberMessage}, nil
	}
}

func (b *Bot) waitLeaderboardNumberMessage(message *tgbotapi.Message) ([]string, error) {
	currentState := b.states[message.From.ID]
	if checkNumber(message.Text) {
		err := currentState.Event(eventGetNumber)
		if err != nil {
			return []string{}, err
		}
		users, err := b.db.GetTopUsers("leaderboard", 5)
		if err != nil {
			return []string{}, err
		}
		var result string
		for i, user := range users {
			result += fmt.Sprintf("%v. %v | %v | %v\n", i+1, user.Username,
				user.Points, ranks.GetRank(user.Points))
		}
		return []string{result}, nil
	} else {
		return []string{wrongNumberMessage}, nil
	}
}

func (b *Bot) waitTaskAnswerMessage(message *tgbotapi.Message) ([]string, error) {
	currentState := b.states[message.From.ID]
	err := currentState.Event(eventGetAnswer)
	var verdict []string
	if err != nil {
		return []string{}, err
	}
	if message.Text == tasks.GetAnswer(currentState.task, currentState.number) {
		err = b.db.IncreaseValue(fmt.Sprintf("leaderboard%v", currentState.task),
			"count_correct", message.From.ID, 1)
		if err != nil {
			return []string{}, err
		}
		err = b.db.IncreaseValue(fmt.Sprintf("leaderboard%v", currentState.task),
			"points", message.From.ID, TasksWeights[currentState.task][0])
		if err != nil {
			return []string{}, err
		}
		verdict = append(verdict, correctAnswerMessage)
	} else {
		err = b.db.IncreaseValue(fmt.Sprintf("leaderboard%v", currentState.task),
			"count_false", message.From.ID, 1)
		if err != nil {
			return []string{}, err
		}
		err = b.db.IncreaseValue(fmt.Sprintf("leaderboard%v", currentState.task),
			"points", message.From.ID, -TasksWeights[currentState.task][1])
		if err != nil {
			return []string{}, err
		}
		verdict = append(verdict, wrongAnswerMessage+tasks.GetAnswer(currentState.task, currentState.number))
		verdict = append(verdict, tasks.GetExplanation(currentState.task, currentState.number))
	}
	return verdict, nil
}

func (b *Bot) waitTrainAnswerMessage(message *tgbotapi.Message) ([]string, error) {
	currentState := b.states[message.From.ID]
	var verdict []string
	if message.Text == tasks.GetAnswer(currentState.task, currentState.number) {
		err := b.db.IncreaseValue(fmt.Sprintf("leaderboard%v", currentState.task),
			"count_correct", message.From.ID, 1)
		if err != nil {
			return []string{}, err
		}
		err = b.db.IncreaseValue(fmt.Sprintf("leaderboard%v", currentState.task),
			"points", message.From.ID, TasksWeights[currentState.task][0])
		if err != nil {
			return []string{}, err
		}
		verdict = append(verdict, correctAnswerMessage)
	} else {
		err := b.db.IncreaseValue(fmt.Sprintf("leaderboard%v", currentState.task),
			"count_false", message.From.ID, 1)
		if err != nil {
			return []string{}, err
		}
		err = b.db.IncreaseValue(fmt.Sprintf("leaderboard%v", currentState.task),
			"points", message.From.ID, -TasksWeights[currentState.task][1])
		if err != nil {
			return []string{}, err
		}
		verdict = append(verdict, wrongAnswerMessage+tasks.GetAnswer(currentState.task, currentState.number))
		verdict = append(verdict, tasks.GetExplanation(currentState.task, currentState.number))
	}

	currentState.number = tasks.GetRandomTask(currentState.task)
	verdict = append(verdict, tasks.GetStatement(currentState.task, currentState.number))
	return verdict, nil
}
