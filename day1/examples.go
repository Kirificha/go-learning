package main

import "fmt"

func iff() {
	x := 5

	if x == 5 {
		fmt.Println("Верно")
	} else {
		fmt.Println("Не верно")
	}

	switch any(x).(type) {
	case string:
		fmt.Println("x это string - значение")
	case int:
		fmt.Println("x это int - значение")
	default:
		fmt.Println("x это default - значение")

		y := 4

		if x == 5 || y == 4 {
			fmt.Println("Всё верно")
		} else {
			fmt.Println("Не верно")
		}

		if x == 5 && y == 3 {
			fmt.Println("Всё нормально")
		}
	}
}

func forrr() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
