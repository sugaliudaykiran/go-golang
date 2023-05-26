package main
import("fmt")
// use the "panic statement" to "immediately end the execution of the program". 
// If our program reaches a point where it cannot be recovered due to some major errors, it's best to use panic.

func divisible(x, y int) int{
    if y==0{
        panic("cannot divisible any number by zero")
    }else{
        return x/y    
    }
}

func main(){
    fmt.Println("Help..! something bad happend..")
    panic("Ending the program..")  // use "go" to avoid the "panic incase"
    fmt.Println("Waiting to execute")
    
    fmt.Println(divisible(11,2))
    fmt.Println(divisible(1,0))
    fmt.Println(divisible(22,2))
}