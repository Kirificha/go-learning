package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	fmt.Println(a)
	b := a
	b[0] = 10
	fmt.Println(a, b)
	c := append(b, 4)
	fmt.Println(len(b), len(c))
	c[0] = 99
	// значение в a и b не изменилось т.к append
	// вышел за границы и создал новый массив "c"
	fmt.Println(b, c)
	fmt.Println(app(a))

	x := make([]int, 2, 4)
	x[0] = 1
	x[1] = 2
	y := x
	z := append(y, 3)
	fmt.Println(len(y), len(z))
	z[0] = 99
	fmt.Println(x, y, z)
	// x и y не видят изменений среза/слайса
	// потому что их длина меньше длины z. если
	// увеличить len x с 2 до 3 - изменения будут видны в отображении через слайс x,y
}

func app(a []int) []int {
	a = append(a, 5)
	return a
}

/* вывод: append с переполнением cap создаст новый массив,
исходный слайс не изменится. Если cap хватает, append запишет в тот же массив,
но длина исходного не изменится, поэтому через исходный слайс новый элемент не виден,
хотя в памяти он есть */
