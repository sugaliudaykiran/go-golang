package main
import (
    "fmt"
    "math"
    "strings"
    "week1/go_packages/calculator"
)

func main(){
    var fname string
    fmt.Print("Enter the fname: ")
    fmt.Scan(&fname)
    Lower := strings.ToLower(fname)
    fmt.Println("Converting into lower: ",Lower)
    
    var num1,num2 float64 = 122, 31
    fmt.Println(math.Max(num1,num2))

    x1, x2 := 10, 21 // imported the calculator package in our program and used its functions.
    fmt.Println(calculator.Adding(x1, x2))
    fmt.Println(calculator.Subtracting(x1, x2))

}

/*

    Note: This file doesn't contain the main package. Hence, the Go compiler doesn't consider this as 
    an executable program and it is created for the sole purpose of sharing and reusing.

*/ 