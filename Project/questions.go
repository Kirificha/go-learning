package main

import "fmt"

func read() {
	var height, weight, age, next int
	var name string

	start := func() {
		fmt.Println("Как вас зовут?")
		fmt.Scan(&name)
		fmt.Scan()
		fmt.Println("Напишите ваш рост, вес, возраст")
		fmt.Scan(&height)
		fmt.Scan(&weight)
		fmt.Scan(&age)
		fmt.Printf("Рост: %d\nВес: %d\nВозраст: %d\nВсё верно?\n", height, weight, age)
		fmt.Println("1 - Да, 2 - Нет")
		fmt.Scan()
		fmt.Scan(&next)
	}

	for {
		start()

		if next == 1 {
			break
		}
	}
}

func Questions(name string) {
	var next int
	fmt.Printf("Приветствую %s! Хотите пройти общий опрос или выбрать категорию?", name)
	fmt.Println("1 - общий опрос, 2 - выбрать категорию")
	fmt.Scan(&next)

loop:
	for {
		switch next {
		case 1:
			runFullSurvey()
			break loop
		case 2:
			runSelectedSurvey(questions)
			break loop
		default:
			fmt.Println("Введите число 1 или 2")
			fmt.Scan(&next)
		}
	}

}

func runFullSurvey() []Scores {
	result := []Scores{
		{Category: "Sleep", Score: 0},
		{Category: "Energy", Score: 0},
	}
	for _, question := range questions {
		num := 0
		fmt.Println(question, "\n 1 - Да, 2 - Нет")
		fmt.Scanln(&num)
		if num == 1 {
			for i := range result {
				if question.Category == result[i].Category {
					result[i].Score += question.Score
				}
			}

		}
	}

	return result
}

// Choice: 1 - Энергия, 2 - Сон
func runSelectedSurvey(questions []Question) []Scores {
	var category string
	score := 0
	num := 0

loop:
	for {
		fmt.Println("Выберите категорию:\n 1 - Сон, 2 - Энергия")
		fmt.Scanln(&num)

		switch num {
		case 1:
			category = "Sleep"
			break loop
		case 2:
			category = "Energy"
			break loop
		default:
			fmt.Println("Введите число отвечающее нужной категории")

		}
	}

	for _, question := range questions {
		if question.Category == category {

			answer := 0
			fmt.Println(question, "\n1 - Да, 2 - Нет")
			fmt.Scanln(&answer)
			if answer == 1 {
				score += question.Score
			}

		}

	}
	result := []Scores{
		{Category: category, Score: score},
	}
	return result
}
