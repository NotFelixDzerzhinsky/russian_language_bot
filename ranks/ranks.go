package ranks

import "fmt"

var ranks = map[int]string{
	-100: "Очень слабый чел",
	-15:  "Слабенкий",
	0:    "Новичок",
	20:   "Я русский",
	35:   "5 по русскому",
	52:   "Культурный петербуржец",
	53:   "Участник школьного этапа по русскому языку",
	75:   "Призёр ОГЭ по русскому языку",
	100:  "Победитель региона по русскому языку",
	150:  "Победитель ВСОШ по русскому",
	250:  "Межнар по русскому",
	500:  "Абсолютный победитель русского медвежонка 2015 года",
	1000: "Потомок древних русов",
}

func GetRanksTable() string {
	var result string
	for val, rank := range ranks {
		if val == 52 {
			continue
		}
		result += fmt.Sprintf("от %d очков %s\n", val, rank)
	}
	return result
}

func GetRank(points int) string {
	last := ranks[-100]
	for val, rank := range ranks {
		if val == points {
			return rank
		}
		if val > points {
			return last
		}
		last = rank
	}
	return last
}
