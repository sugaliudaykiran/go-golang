package main
import (
    "fmt"
    _ "errors"
)
//  error occurs, the execution of a program stops completely --> errors.New(), fmt.Errorf()
func divisible(x, y int) (int,error){
    if y==0{
        // return 0,errors.New("Can't be divisible by anything by zero(0)")
        return 0, fmt.Errorf("Can't be divisible by anything by zero(%d)",y)
    }else{
        return x/y, nil
    }
}

func main(){
    res,err := divisible(11, 0)
    if err!=nil{
        fmt.Println(err)
    }else{
        fmt.Printf("divisible of x/y: %d", res)
    }
}


/*
	When an error occurs, the execution of a program stops completely with a built-in error message. 
	In Go, we . Hence, it is important to handle those exceptions in our program.

Unlike other programming languages, we don't use try/catch to handle errors in Go. We can handle errors using:

		New() Function
		Errof() Function

*/ 