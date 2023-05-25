package go_packages

// declaring a variable in the cal, can be only used in go_packages
var name string = "uday"

// creating add function
func Adding(x, y int)int{
	return x+y
}
// creating subtract function
func Subtracting(x, y int)int{
	return x-y
}

/*
   Note: This file doesn't contain the main package.
    Hence, the Go compiler doesn't consider this as an executable 
	program and it is created for the sole purpose of sharing and reusing.
*/ 