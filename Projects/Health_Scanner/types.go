package main

type User struct {
	Name          string
	Age           int
	Height        float64
	Weight        float64
	ActivityLevel string
}

type Scores struct {
	Category string
	Score    int
}

type Spheres struct {
	Category string
	Name     string
	Scoremin int
	//я тут думаю записывать все такие сферы жизни.
	//Тут копятся баллы набранные пользователем, чем больше - тем хуже
}

type Result struct {
	Category string
	Name     string
}

type Question struct {
	Text     string
	Category string
	Score    int
}

// А вот здесь эти баллы сверяются уже и выдают результат Insomnia/BadSleep e.t.c
var results = []Spheres{
	{Category: "Sleep", Name: "insomnia", Scoremin: 85},
	{Category: "Sleep", Name: "bad sleep", Scoremin: 60},
	{Category: "Sleep", Name: "common sleep", Scoremin: 35},
	{Category: "Sleep", Name: "good sleep", Scoremin: 0},
}

var questions = []Question{
	{Text: "Вы спите менее 7-и часов?", Category: "Sleep", Score: 10},
	{Text: "Часто ли вы просыпаетесь ночью?", Category: "Sleep", Score: 10},
	{Text: "Бывает ли у вас бессоница или трудности с засыпанием?", Category: "Sleep", Score: 10},
	{Text: "Проводите ли вы время перед сном в телефоне/компьютере?", Category: "Sleep", Score: 10},
	{Text: "Чувствуете ли вы себя как в тумане в течение дня?", Category: "Energy", Score: 10},
	{Text: "Ломит ли вас в середине дня, хочется полежать, или нет такого?", Category: "Energy", Score: 10},
}

func Condition(scores []Scores, spheres []Spheres) ([]Result, error) {
	var result []Result
	for _, score := range scores {
		for _, sphere := range spheres {
			if score.Category == sphere.Category {
				if score.Score >= sphere.Scoremin {
					add := Result{Category: sphere.Category, Name: sphere.Name}
					result = append(result, add)
					break

				}

			}
		}
	}
	return result, nil
}
