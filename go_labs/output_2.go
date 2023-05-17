package main
import ("fmt")

func main(){
    // verbs can be used with all data types:
    var var1,var2 = "kiran", 513
    fmt.Printf("val of var1: %v and var2: %v \n", var1, var2)
    fmt.Printf("val of var1: %#v and var2: %#v \n", var1, var2)
    fmt.Printf("type of var1: %T and var2: %T \n", var1, var2)
    var per = 71
    fmt.Printf("val of per: %v%% \n",per)

	// verbs can be used with the integer data type:
	var x, y = 10, 15
    fmt.Printf("base 2 of %v: %b and base 2 of %v: %b \n",x,x,y,y)
    fmt.Printf("base 10 of %v: %d and base 10 of %v: %d \n\n",x,x,y,y)
    
    fmt.Printf("lower case char of %v: %x and lower case char of %v: %x \n",x,x,y,y)
    fmt.Printf("upper case char of %v: %X and upper case char of %v: %X \n\n",x,x,y,y)
    
	// verbs can be used with the strign data type:
    var str1, str2 = "uday", "kiran"
    fmt.Printf("value of str1: %s and value of str2: %s \n",str1,str2)
    fmt.Printf("value of str1: %q and value of str1: %q \n\n",str1,str2)
    fmt.Printf("value of str1: % x and value of str2: % x \n",str1,str2) // hex value of each char
	
	// verbs can be used with the float data type:
	var b1, b2 bool = true, false // %t -> %v
    fmt.Printf("value of b1: %t and value of b2: %v \n",b1,b2)
    var f1, f2 = 3.14, 4.02838
    fmt.Printf("value of f1: %e and value of f2: %f \n",f1,f2)
    fmt.Printf("value of f1: %.4f and value of f2: %.2f \n",f1,f2)
}