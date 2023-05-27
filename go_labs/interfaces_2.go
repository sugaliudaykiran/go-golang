package main

import (
	"fmt"
	// 	"math"
)

type size interface {
	max() float64
}

type person struct {
	height int
	weight float64
}

type animal struct {
	height int
	weight float64
}

func (p *person) max() float64 { // here argument are passing as an Pointers..
	return float64(float64(p.height*2) + p.weight)
}

func (a animal) max() float64 {
	return float64(float64(a.height*2) + a.weight)
}

func main() {
	p1 := person{11, 3.2}
	a1 := animal{22, 42.1}
	maxWeight := []size{&p1, a1} // dereferencing the
	for _, v := range maxWeight {
		fmt.Println(v.max())
	}
}
