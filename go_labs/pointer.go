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

func main() {
	x := 5
	zero(x) // Here x is value of zero after the function call also.
	fmt.Printf("x: %v \n", x)
	zero2(&x)
	fmt.Printf("x: %v \n", x)

	yptr := new(int)
	one(yptr)
	fmt.Println(*yptr) // garbage collected programming language which means memory is cleaned up automatically when nothing refers to it anymore.

}
