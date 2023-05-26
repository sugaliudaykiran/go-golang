package main
import (
    // "strings"
    "fmt"
    "strconv"
    // "math"
)

func main(){
    // type casting
    var1,var2,var3 := 121,212.21,21
    fmt.Printf("converting var1 - int to float: %.2f \n",float64(var1)/float64(var3))
    fmt.Printf("converting var2 - float to int: %d \n",int64(var2))
    
    // converting string to int vi versa
    x1,x2 := 21,"414"
    fmt.Printf("converting int - string: %s \n", strconv.Itoa(x1))
    res ,_ := strconv.Atoi(x2)
    fmt.Printf("converting string - int: %d \n",res)
}
/*
    Type conversion is a great way to change the original data type. But unfortunately, in some cases, it loses precision.
         It should be done sparingly. The needs of type conversion lie in the fact Go doesn’t support 
        implicit type conversions. So in many cases, it should be done separately.
    

    Note: As Golang has a strong type system, it doesn’t allow to mix(like addition, subtraction, multiplication, division, etc.) 
    the numeric types in the expressions and also you are 
        not allowed to perform an assignment between the two mixed types.
*/ 
