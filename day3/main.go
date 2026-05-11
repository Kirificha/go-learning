package main

import (
	"fmt"
)

type speaker interface {
	speak() string
}

type animal struct {
	name  string
	sound string
}

type Cat struct {
	name  string
	sound string
}

type Dog struct {
	name  string
	sound string
}

func (an *animal) speak() string {
	return an.name + " говорит " + an.sound
}

func (d *Dog) speak() string {
	return d.name + " говорит " + d.sound
}

func (c *Cat) speak() string {
	return c.name + " говорит " + c.sound
}

func main() {

	var s speaker
	s = &Cat{name: "Мурзик", sound: "мяу"}
	fmt.Println(s.speak())
	s = &Dog{name: "Барс", sound: "гав"}
	fmt.Println(s.speak())

	r := rectangle{width: 10, height: 20}
	fmt.Println("Area:", r.area(), "Perimeter:", r.perim())

}
