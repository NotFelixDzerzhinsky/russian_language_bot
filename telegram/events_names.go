package telegram

var TasksNumbers = []int{4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22}

var TasksWeights = [][]int{
	{0, 0}, // 0
	{0, 0}, // 1
	{0, 0}, // 2
	{0, 0}, // 3
	{1, 1}, // 4
	{1, 1}, // 5
	{1, 1}, // 6
	{1, 1}, // 7
	{2, 1}, // 8
	{3, 1}, // 9
	{2, 1}, // 10
	{2, 1}, // 11
	{1, 1}, // 12
	{1, 1}, // 13
	{2, 1}, // 14
	{2, 1}, // 15
	{1, 1}, // 16
	{2, 1}, // 17
	{2, 1}, // 18
	{2, 1}, // 19
	{2, 1}, // 20
	{1, 1}, // 21
	{1, 1}} // 22

// commands
const commandStart = "start"
const commandHelp = "help"
const commandExit = "exit"
const commandTask = "task"
const commandTrain = "train"
const commandLeaderboard = "leaderboard"
const commandRanks = "ranks"

// states
const stateDefault = "default"
const stateWaitTaskNumber = "wait_task_number"
const stateWaitTaskAnswer = "wait_task_answer"
const stateWaitLeaderboardNumber = "wait_leaderboard_number"
const stateWaitTrainNumber = "wait_train_number"
const stateWaitTrainAnswer = "wait_train_answer"

// events
const eventCommandTask = "command_task"
const eventCommandExit = "command_exit"
const eventCommandTrain = "command_train"
const eventCommandLeaderboard = "command_leaderboard"
const eventGetNumber = "get_task_number"
const eventGetAnswer = "get_task_answer"

// messages
// Главное приветственное сообщение, /start, /help
const startMessage = "Всееееем привет мы здесь сделаем классное ааа todo короче" // todo

// Когда дают несуществующую команду
const wrongCommandMessage = "Я не знаю, что ты себе придумал... но такой команды у меня нет, прости"

// Когда дают ненужную команду в данный момент времени
const unwantedCommandMessage = "Вот зачем вот здесь ты пишешь эту команду? Не время для этого! Лучше ответь на прошлый вопрос."

// Когда дают ненужный текст в данный момент времени
const unwantedTextMessage = "Что ты такое пишешь??? Напиши то, что мне нужно!"

// Если пишут какой-то текст в stateDefault
const defaultTextMessage = "Что тебе нужно? Чтобы я работал напиши какую-то нибудь команду, если что-то непонятно напиши /help"

// Когда дают неверный номер задач
const wrongNumberMessage = "Увы, у меня нет задания с таким номером, попробуй ещё разок"

// Неверный ответ
const wrongAnswerMessage = "О неееет, ты даже не представляешь как же ты ошибаешься... Правильный ответ это:"

// Верный ответ
const correctAnswerMessage = "Хароооооош!"

// Просит номер задачи у пользователя
const getNumberMessage = "Отличненько, теперь тебе нужно ввести номер задачи из данных ниже."

// Когда человек пишет exit
const exitCommandMessage = "Ну нет так нет..."
