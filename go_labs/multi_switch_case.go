package main
import ("fmt")

func main(){
    var Num int  // It is possible to have multiple values for each case in the switch statement:
    fmt.Print("Plese Enter the Num: ")
    fmt.Scanf("%v \n",&Num)
    switch Num {
        case 1, 3, 5:
          fmt.Println("Num is an Odd Number")
        case 2, 4:
          fmt.Println("Num is an Even Number")
        default:
          fmt.Println("Pls Enter the day in btw 1 to 6")
    }
}