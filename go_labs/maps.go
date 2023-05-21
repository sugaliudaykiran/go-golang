package main
import ("fmt")

func main(){
    // var mp = map[string]int{"one":1, "two":2, "three":3}
    mp := map[string]int{"one":1, "two":2, "three":3}
    // var mp = make(map[string]int)  --> Used to create a map of empty.. (below one also)
    // mp := make(map[string]int)   
    // var mp map[string]int   --> avoid to create in this format bcz runtime panic
    
    mp["one"]=10; mp["two"]=200
    fmt.Println(mp)
    fmt.Println(mp["one"])
    mp["two"] = 20
    fmt.Println(mp["two"])
    
    // removing element by using delete()
    delete(mp, "three")
    fmt.Println(mp)
    
    // identifing element is "present or not" in map, skiping items by using "_", like val,_ (or) _,existing
    val,existing := mp["two"]
    fmt.Println(val, existing)
    
    // refering the hash table
    h1 := map[string]int{"a":1,"b":2,"c":3}
    h2 := h1
    
    // changing the h2,  will effect the h1 bcz there are refering to same hash table
    h2["b"]=33
    fmt.Println(h1["b"], h2["b"])
    
    // Iterating using range over "map"
    for k,v := range h1{
        fmt.Println(k," : ",v)
    }
    
    // Iterate Over Maps in a Specific Order
    m1 := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}
    myslice := []string{}
    myslice = append(myslice, "one", "two", "three", "four")  // mention the Order
    
    for k,v := range m1{
        fmt.Println(k,": ",v)
    }
    fmt.Println()
    for _,element := range myslice{
        fmt.Println(element,": ",m1[element])
    }
    
}

/*

Maps are used to store data values in key:value pairs.

Each element in a map is a key:value pair.

A map is an unordered and changeable collection that does not allow duplicates.

The length of a map is the number of its elements. You can find it using the len() function.

The default value of a map is nil.

Maps hold references to an underlying hash table.

Go has multiple ways for creating maps.
======================================================
Allowed Key Types
The map key can be of any data type for which the equality operator (==) is defined. These include:

Booleans
Numbers
Strings
Arrays
Pointers
Structs
Interfaces (as long as the dynamic type supports equality)
Invalid key types are:

Slices
Maps
Functions
These types are invalid because the equality operator (==) is not defined for them.

================================
Check For Specific Elements in a Map
You can check if a certain key exists in a map using:

Syntax
val, ok :=map_name[key]
If you only want to check the existence of a certain key, you can use the blank identifier (_) in place of val.


===========================
Maps Are References
Maps are references to hash tables.

If two map variables refer to the same hash table, changing the content of one variable affect the content of the other.

=======================
Iterating Over Maps
You can use range to iterate over maps.

=================================
Iterate Over Maps in a Specific Order
Maps are unordered data structures. If you need to iterate over a map in a specific order, you must have a separate data structure that specifies that order.
*/ 