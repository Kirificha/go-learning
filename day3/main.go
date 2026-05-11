package main

import "fmt"

func main() {
	a := 2.0
	b := 3.0

	fmt.Println(
		Add(a, b),
		Substract(a, b),
		Multiply(a, b),
	)

	result, err := Divide(a, b)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println(result)
	}
}
