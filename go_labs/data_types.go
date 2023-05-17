package main
import ("fmt")

func main(){
    // data types => int, uint(non negative), float, bool
    var var1, var3, var4 = -100, 10.1312, true
    var var2 uint = 321
    //  var1, var3, var4 := -100, 10.1312, true
    fmt.Printf("values of different data types: \n var1: %v, var2: %v, var3: %v, var4: %v \n\n", var1,var2,var3,var4)
    fmt.Printf("types of different data types: \n var1: %T, var2: %T, var3: %T, var4: %T \n\n", var1,var2,var3,var4)
}