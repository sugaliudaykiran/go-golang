package main

import (
	"fmt"
)

func zero(x int) {
	x = 0
}

func zero2(x *int) {
	*x = 0 // * is also used to “dereference” pointer variables.
}
func one(yptr *int) {
	*yptr = 1
}

// "new" takes a type as an argument, allocates enough memory to fit a value of that type and returns a pointer to it.

type Vertex struct {
    X, Y int
}

func main() {
	x := 5
	zero(x) // Here x is value of zero after the function call also.
	fmt.Printf("x: %v \n", x)
	zero2(&x)
	fmt.Printf("x: %v \n", x)

	yptr := new(int) // Here we were creating a memory address to store a value of integer type
	one(yptr)
	fmt.Println(*yptr) // garbage collected programming language which means memory is cleaned up automatically when nothing refers to it anymore.
    
    p := Vertex{1,2}
    q := &p
    r := &Vertex{1,2}  // q,r are pointers to a vertex
    
    var s *Vertex = new(Vertex)  // create ptr to new struct instance
    fmt.Println(p,q,r,s)
}
