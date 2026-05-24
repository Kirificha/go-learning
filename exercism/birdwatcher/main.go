package main

import "fmt"

func main() {
	birdsPerDay := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
	TotalBirdCount(birdsPerDay)
	BirdsInWeek(birdsPerDay, 2)
	FixBirdCountLog(birdsPerDay)
	FixBirdCountLog2(birdsPerDay)
	FixBirdCountLog3(birdsPerDay)
	FixBirdCountLog4(birdsPerDay)
}

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	sum := 0
	for i := range birdsPerDay {
		sum += birdsPerDay[i]
	}
	fmt.Println(sum)
	return sum
	panic("Please implement the TotalBirdCount() function")
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	sum := 0
	for i := 0; i < 7; i++ {
		d := 7*week - 7 + i
		sum += birdsPerDay[d]

		fmt.Println(sum)
	}
	return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i]++
	}
	fmt.Println(birdsPerDay)
}

func FixBirdCountLog2(birdsPerDay []int) {
	for i := range birdsPerDay {
		birdsPerDay[i*2]++
	}
	fmt.Println(birdsPerDay)
}

func FixBirdCountLog3(birdsPerDay []int) {
	for key, value := range birdsPerDay {
		birdsPerDay[key*2] = value + 1
	}

}

func FixBirdCountLog4(birdsPerDay []int) {
	for _, value := range birdsPerDay {

		birdsPerDay[value] = value + 1
	}
}
