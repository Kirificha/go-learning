package main

import "fmt"

func main() {
	fizz()
}

func fizz() {

	for num := 1; num <= 30; num++ {
		if num%3 == 0 && num%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if num%3 == 0 {
			fmt.Println("Fizz")
		} else if num%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(num)
		}
	}
}

func sum() {
	summ := 0
	for num := 1; num <= 100; num++ {
		summ += num
	}
	fmt.Println(summ)
}

func days() {
	i := 4
	fmt.Println("Сегодня")
	switch i {
	case 1:
		fmt.Println("Понедельник")
	case 2:
		fmt.Println("вт")
	case 3:
		fmt.Println("ср")
	case 4:
		fmt.Println("чт")
	case 5:
		fmt.Println("пт")
	case 6:
		fmt.Println("сб")
	case 7:
		fmt.Println("вс")
	}
}
