package main

type rect struct {
	width, height int
}

func main() {
	r := rect{width: 10, height: 5}
	println("area: ", r.area())
	println("perim: ", r.perim())

	rp := &r
	println("area: ", rp.area())
	println("perim: ", rp.perim())
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

/* OUTPUT
area:  50
perim:  30
area:  50
perim:  30
*/