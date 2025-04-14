package intermediate

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rectangle struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perim() float64 {
	return (r.height + r.width) * 2
}

func (c circle) area() float64 {
	return math.Pi * (c.radius * c.radius)
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) diameter() float64 {
	return c.radius * 2
}

// Test
type rectangleOne struct {
	width  float64
	height float64
}

func (r1 rectangleOne) area() float64 {
	return r1.height * r1.width
}

func measure(g geometry) {
	fmt.Println("Geometry:", g)
	fmt.Println("Geometry area:", g.area())
	fmt.Println("Geometry perimeter:", g.perim())
}

func main() {
	r := rectangle{
		width:  10.25,
		height: 20.32,
	}

	c := circle{
		radius: 12.69,
	}

	r1 := rectangleOne{
		width:  10,
		height: 20,
	}
	_ = r1

	measure(r)
	measure(c)
	// measure(r1) ERROR!!

	myPrinter(1, 2, 3, "Satoshi", 1.5)
}

func myPrinter(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v)
	}
}
