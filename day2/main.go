package main

import "fmt"

func main() {
	defer fmt.Println("done")
	var nums = [5]int{-2990, 2, 3, 4, 5}
	min, max := minMax(nums[:])
	fmt.Println(min, max)
}
