package main
import ("fmt")

func main(){
    ar := [4]int{1, 2,3 , 4}
    fmt.Printf("Values of ar: %v \n",ar)
    // myslice := []string{}  creating empty slice
    // myslice := []int{11, 21, 32, 31}
    // myslice := ar[1:3]  // Not including the end item of it..
    // myslice := ar[:3]
    myslice := ar[1:]   
    fmt.Printf("Values of myslice: %v \n",myslice)
    fmt.Printf("Length of myslice: %v, Capacity of myslice: %v \n",len(myslice), cap(myslice))
    
    slice2 := []string{"uday","kiran","sugali"}
    fmt.Printf("string of slice2: %v \n", slice2)
    
    slice3 := make([]bool,4)
    fmt.Printf("using make for creating slice: %v \n", slice3)
    
    fmt.Println("first value of slice2: ",slice2[0],"\t Length of slice2: ",len(slice2))
    
    slice3[0] = true
    fmt.Println("0 th Index of slice3: ", slice3[0])
    
    // Append Elements To a Slice
    myslice1 := []int{11, 21, 32, 41}
    myslice2 := []int{0,1, 2, 3}
    
    // appending to itself..
    myslice1 = append(myslice1, 51, 61, 71)
    fmt.Println(myslice1, myslice2)
    
    // appending from one slice to other
    myslice3 := append(myslice1,myslice2...)  // ... is mandatory while using append on two slice..
    fmt.Println("values of myslice3: ",myslice3)
    
    // Using copy() in-case of "Memory Efficiency"
    cur_slice := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
    fmt.Println()
    fetch_slice := cur_slice[0:5]
    temp_slice := make([]int,len(fetch_slice))
    fmt.Println("before copy final_slice: ", temp_slice)
    copy(temp_slice, fetch_slice)
    
    fmt.Println("final_slice: ", temp_slice)
    
}