package main

import "fmt"

func main() {
	r := rect{5, 5}

	measure(r)

}

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	with, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return 2*r.height + 2*r.with
}

func (r rect) perim() float64 {
	return 2*r.with + 2*r.height
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
