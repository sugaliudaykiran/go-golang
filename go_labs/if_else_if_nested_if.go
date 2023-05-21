package main
import ("fmt")

func main(){
    //  year is leap or not
    var year int
    fmt.Scanf("%v\n", &year)
    fmt.Println("Year: ",year)
    
    if year%4==0{
        if  year%100!=0 || year%400==0 {
            fmt.Println("yes")
        }
    }else{
        fmt.Println("no")
    }
    
    x1, x2 := 10, 21
    if x1 > x2{
        fmt.Println("x1 is greater than x2")
    }else if x1 < x2{
        fmt.Println("x2 is greater than x1")
    }else{
        fmt.Println("x1 is equal to x2")
    }
}