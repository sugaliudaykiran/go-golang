package main
import ("fmt")

func main(){
    //  Arrays
    var arr1 = [4]int{11, 23, 32, 11};
    var arr2 = [...]float64{1.12, 0.22, 9.3, 4.1, 2.5};
    
    arr3 := [3]string{"Hello", "World", "Hii"}
    arr4 := [...]bool{true, false}
    
    fmt.Printf("values of arr1: %v \n", arr1)
    fmt.Printf("values of arr2: %v \n", arr2)
    fmt.Printf("values of arr3: %v \n", arr3)
    fmt.Printf("values of arr4: %v \n", arr4)
    
    fmt.Println(arr3[2])
    arr3[2] = "Uday"
    fmt.Println(arr3[2])
    
    // Initialization only specific element in an array - using indexing
    arr5 := [...]int{1:10, 2:20, 4:40}
    fmt.Println(arr5)
    
    // length of array
    var arrLen = len(arr5)
    fmt.Println(arrLen)
    
    
    /*
    
    var array_name = [length]datatype{values} // here length is defined (or) 
    var array_name = [...]datatype{values} // here length is inferred
 
    array_name := [length]datatype{values} // here length is defined (or)
    array_name := [...]datatype{values} // here length is inferred
    
    */ 
}