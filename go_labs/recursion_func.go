package main
import ("fmt")

func hello(fname string){
    fmt.Println("Hello " + fname)
}

func testcount(x int) int {
    if x==11{
        return 0
    }else{
        fmt.Println(x)
        return testcount(x+1)
    }
}

func factorial(n int) (res int){
    if n>0{
        res=n*factorial(n-1)
    }else{
        res=1
    }
    return
}

func main(){
    hello("uday")
    testcount(1)
    fmt.Println(factorial(5))
}