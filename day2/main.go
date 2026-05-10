package main

import "fmt"

func main() {
	defer fmt.Println("done")
	var names = [5]string{"Виталя", "Саня", "Кирилл", "Игорь", "Мужик"}
	result := slices(names[:])
	fmt.Println(result)
}
