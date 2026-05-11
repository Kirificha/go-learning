package main

type rectangle struct {
	width, height int
}

func (r *rectangle) area() int {
	return r.height * r.width
}

func (r *rectangle) perim() int {
	return r.height*2 + r.width*2
}
