package main

import (
	"fmt"
)

func myfun() {
	fmt.Println("I just got executed!")
}

func myfun3(fname string) {
	fmt.Println("Welcome ", fname)
}

func myfun2(x int, y int) {
	if x > y {
		fmt.Println("X is Greater than Y")
	} else if x < y {
		fmt.Println("Y is Greater than X")
	} else {
		fmt.Println("X and Y are equal")
	}
}

// return - should mention "type and return"
func sumof2(x int, y int) int {
	return x + y
}

// return by variable
func sumof3(x int, y int, z int) (result int) {
	result = x + y + z

	return
}

// return - for multiple values are returing
func sumof_2(x int, y string) (int, string) {
	return x + x, "Hello " + y
}

// return by variable of multiple variable
func sumof_3(x int, y string) (res int, txt string) {
	res = x + x
	txt = "Hello " + y

	return
}

func main() {
	myfun()
	myfun3("uday")
	myfun2(10, 21)

	fmt.Println(sumof2(1, 3))
	result := sumof2(2, 3)
	fmt.Println(result)

	fmt.Println(sumof3(1, 2, 4))
	res := sumof3(1, 2, 3)
	fmt.Println(res)

	fmt.Println(sumof_2(1, "uday"))
	res, txt := sumof_2(2, "kiran")
	fmt.Println(res, txt)

	fmt.Println(sumof_3(2, "sugali"))
	res2, txt2 := sumof_3(2, "sai")
	fmt.Println(res2, txt2)

	// If we (for some reason) do not want to use some of the returned values, we can add an underscore (_), to "omit" this value.
	_, txt3 := sumof_3(2, "kumar")
	fmt.Println(txt3)

	x, _ := sumof_3(1, "ram")
	fmt.Println(x)

}
