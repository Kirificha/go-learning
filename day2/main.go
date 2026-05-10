package main

import "fmt"

func main() {
	defer fmt.Println("done")
	text := "Супер дупер супер текст придумал супер"
	fmt.Println(doubles(text))

	box := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(chet(box))
}
