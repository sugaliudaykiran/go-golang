package main
import ("fmt")

func main(){
    // A type assertion provides access to an interface value's underlying concrete value.
    var i interface{} = "hello"
    
    value1:=i.(string)
    
    fmt.Println(value1)
    
    val2, ok:=i.(int)  // return two values: the underlying value and a boolean value that reports whether the assertion succeeded.
    if ok{
        fmt.Println("val2 : ", val2)
    }else{
        fmt.Printf("Invalid, type-%T doesn't exist.", val2)
    }
    val3:=i.(float64)
    fmt.Println(val3) // panic
}


/*
    
    Type assertions
A type assertion provides access to an interface value's underlying concrete value.

t := i.(T)
This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.

If i does not hold a T, the statement will trigger a panic.

To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.

t, ok := i.(T)
If i holds a T, then t will be the underlying value and ok will be true.

If not, ok will be false and t will be the zero value of type T, and no panic occurs.

Note the similarity between this syntax and that of reading from a map.

*/ 