package telegram

import (
	"context"
	"github.com/looplab/fsm"
)

// StateMachine states - машина состояний, number - номер упражнения в данный момент,
// -1 если нет, task - номер задания в таблице, -1 если нет
type StateMachine struct {
	states *fsm.FSM
	number int
	task   int
}

func NewStateMachine() *StateMachine {
	return &StateMachine{states: fsm.NewFSM(
		stateDefault,
		fsm.Events{
			{Name: eventCommandTask, Src: []string{stateDefault}, Dst: stateWaitTaskNumber},
			{Name: eventCommandTrain, Src: []string{stateDefault}, Dst: stateWaitTrainNumber},
			{Name: eventCommandLeaderboard, Src: []string{stateDefault}, Dst: stateWaitLeaderboardNumber},
			{Name: eventGetNumber, Src: []string{stateWaitTaskNumber}, Dst: stateWaitTaskAnswer},
			{Name: eventGetNumber, Src: []string{stateWaitLeaderboardNumber}, Dst: stateDefault},
			{Name: eventGetNumber, Src: []string{stateWaitTrainNumber}, Dst: stateWaitTrainAnswer},
			{Name: eventGetAnswer, Src: []string{stateWaitTaskAnswer}, Dst: stateDefault},
			{Name: eventCommandExit, Src: []string{stateWaitTaskAnswer}, Dst: stateDefault},
			{Name: eventCommandExit, Src: []string{stateWaitTaskNumber}, Dst: stateDefault},
			{Name: eventCommandExit, Src: []string{stateWaitTrainNumber}, Dst: stateDefault},
			{Name: eventCommandExit, Src: []string{stateWaitTaskAnswer}, Dst: stateDefault},
			{Name: eventCommandExit, Src: []string{stateWaitTaskNumber}, Dst: stateDefault},
			{Name: eventCommandExit, Src: []string{stateWaitTrainAnswer}, Dst: stateDefault},
		},
		fsm.Callbacks{},
	),
		number: -1,
		task:   -1,
	}
}

func (s *StateMachine) Current() string {
	return s.states.Current()
}

func (s *StateMachine) Event(event string) error {
	return s.states.Event(context.Background(), event)
}
