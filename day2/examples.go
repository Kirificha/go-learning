package main

import (
	"fmt"
	"strings"
)

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

// задания из дня 3
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Нельзя делить на ноль")
	}
	return a / b, nil

}

func minMax(nums []int) (int, int) {
	minn := nums[0]
	maxx := nums[0]
	for _, v := range nums {
		if v > maxx {
			maxx = v
		}
		if v < minn {
			minn = v
		}
	}
	return minn, maxx
}

// задания из дня 4
func slices(slice []string) []string {
	names := []string{}
	for _, v := range slice {
		if len([]rune(v)) > 4 {
			names = append(names, v)
		}
	}
	return names
}

func doubles(text string) any {
	words := strings.Split(text, " ")
	mapa := map[string]int{}
	for _, v := range words {
		mapa[v]++
	}
	return mapa
}

func chet(nums []int) []int {
	s := []int{}
	for i := range nums {
		if nums[i]%2 == 0 {
			s = append(s, nums[i])
		}
	}
	return s
}
