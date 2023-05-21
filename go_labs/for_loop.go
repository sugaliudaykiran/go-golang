package main
import ("fmt")

func main(){
    // for loop - existing loops in golang
    for i:=0; i<5; i++{
        if i==3{
            continue
        }else if i==5{
            break
        }
        fmt.Println(i)
    }  
    fmt.Println("\n")
    
    //  Nested loop
    arr1 := []string{"1","2"}
    arr2 := []string{"a", "b", "c", "d"}
    for i:=0; i<len(arr1); i++{
        for j:=0; j<len(arr2); j++{
            fmt.Println(arr1[i], arr2[j])
        }
        fmt.Println()
    }
    
    // range keyword is used to more easily iterate over an array, slice or map. It returns both the index and the value.
    myarr := []string{"aa", "bb", "cc", "dd"}
    for ind,val := range myarr{
        fmt.Println(ind, val)
    }
    fmt.Println()
    
    // Tip: To only show the value or the index, you can omit the other output using an underscore (_).
    for _,val := range myarr{ // with "_" declaring the ind
        fmt.Println(val)
    }
    fmt.Println() // with "_" declaring the val
    for ind,_ := range myarr{
        fmt.Println(ind)
    }
    
    fmt.Println()  // without declaring the val
    for ind := range myarr{
        fmt.Println(ind)
    }
    
}