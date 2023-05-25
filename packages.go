package main

import (
	"example.com/packages/go_packages"
	"fmt"
	"math"
	"strings"
)

func main() {
	var fname string
	fmt.Print("Enter the fname: ")
	fmt.Scan(&fname)
	Lower := strings.ToLower(fname)
	fmt.Println("Converting into lower: ", Lower)

	var num1, num2 float64 = 122, 31
	fmt.Println(math.Max(num1, num2))

	x1, x2 := 10, 21 // imported the calculator package in our program and used its functions.
	fmt.Println(go_packages.Adding(x1, x2))
	fmt.Println(go_packages.Subtracting(x1, x2))

	// calling the function of greet
	go_packages.Greet()

	// trying read the variable of "fname"
	fmt.Println(go_packages.name)
	// name not exported by package go_packages (compile)
}

/*

   Note: This file doesn't contain the main package. Hence, the Go compiler doesn't consider this as
   an executable program and it is created for the sole purpose of sharing and reusing.

*/
