package tasks

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
)

type Task struct {
	number      int
	task        int
	statement   string
	answer      string
	explanation string
}

const maxNumber = 30

var tasksList [maxNumber][]Task

func Init(filePath string, task int) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("csv open error: %w", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	result, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("csv read error: %w", err)
	}
	for id, row := range result {
		if row[0] == "statement" {
			continue
		}
		tasksList[task] = append(tasksList[task], Task{number: id - 1, task: task, statement: row[0],
			answer: row[1], explanation: row[2]})
	}
	return nil
}

func GetRandomTask(task int) int {
	number := rand.Intn(len(tasksList[task]))
	return number
}

func GetStatement(task int, number int) string {
	return tasksList[task][number].statement
}

func GetAnswer(task int, number int) string {
	return tasksList[task][number].answer
}

func GetExplanation(task int, number int) string {
	return tasksList[task][number].explanation
}
