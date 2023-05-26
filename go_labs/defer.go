package main
import ("fmt")

func divisible(x,y int)int{
    defer func(){     // Here In this case, error case doesn't effect the flow function calls of "main()"
        fmt.Println(recover())  // with help of defer && recover()
    }()
    return x/y
}

func main(){
    // defer -> to delay the execution of functions that might cause an error.
    defer fmt.Println("one")
    fmt.Println("two") // this defer statements will execute after the execution of all other functions need to done.
    defer fmt.Println("three")
    fmt.Println("four")  // In case of multiple defer statements, the order of execution will be "LIFO"(stack)
    
    // defer inside the function (or) can also use in function calls.
    fmt.Println(divisible(99, 0))
    fmt.Println(divisible(25, 5))
    fmt.Println(divisible(2, 0))
    fmt.Println(divisible(20, 2))
    fmt.Println(divisible(11, 0))
    fmt.Println("---The End---")
}