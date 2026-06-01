package main

import "fmt"

func main() {
	/*	score := []Scores{
			{Category: "Sleep", Score: 61},
			{Category: "Energy", Score: 35},
		}
		i, _ := Condition(score, results)
		fmt.Println(i) */

	read()
	scores := Questions(name)
	result, _ := Condition(scores, results)
	fmt.Println(result)
}
