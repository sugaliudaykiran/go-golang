package main
import ("fmt")

func main(){
    var x, y, z = 1, 2, 3
    fmt.Print(x, y, z,"\n")  // go - output formatting
    var str1, str2 = "Hello", "World"
    fmt.Print(str1," ", str2,"\n")
    
    
    fmt.Println(str2, str1)
    
    
    fmt.Printf("The value of str1 is: %v and type of it is: %T ",str1,str1)
    fmt.Println()
    fmt.Printf("The value of x is: %v and type of it is: %T \n",x,x)
}