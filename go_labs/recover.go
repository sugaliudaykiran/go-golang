package main
import("fmt")
// use the "panic statement" to "immediately end the execution of the program". 
// sometimes it might be important for a program to complete its execution and get some required results.
// "recover statement" prevents the termination of the program and recovers the program from panic.
func handlePanic(){
    a := recover()
    if a!=nil{
        fmt.Println("RECOVER ",a)
    }
}

func divisible(x, y int){
    defer handlePanic() // to excute if panic happends before prg terminates..
    
    if y==0{
        panic("cannot divisible any number by zero")
    }else{
        fmt.Println(x/y)
    }
}

func main(){
    
    divisible(11,2)
    divisible(1,0)
    divisible(22,2)
}

/*

Here, notice two things:
    
    * We are calling the handlePanic() before the occurrence of panic.It's because the program will be terminated
          if it encounters panic and we want to execute handlePanic() before the termination.
    * We are using defer to call handlePanic(). It's because we only want to handle the panic after it occurs, 
          so we are deferring its execution.


*/ 