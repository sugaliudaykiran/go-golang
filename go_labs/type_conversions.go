package main
import (
    "fmt"
)
type person struct {
	name string
	age int
}

type child struct {
	name string
	age int
}


type pet struct{
  name string
}
// type pet child  := "would assign the type of exisiting struct"

//  myInt is a new type who's base type is `int`
type MyInt int
// The Addone method works on `myInt` types, but not regular `int`s
func (i MyInt)Addone() MyInt{
    return i+1
}

func (i MyInt)Addtwo() MyInt{
    return i+2
}

func main() {
	bob := person{
		name: "bob",
		age: 15,
	}
    babyBob := child(bob)
    
    // "babyBob := pet(bob)" would result in a compilation error
	fmt.Println(bob, babyBob)
	// babyCat := pet(bob)
	
    // 	You can declare a new type from a base type, or by creating a composite type:
    var i MyInt = 4
    fmt.Println(i.Addone())
    fmt.Printf("%T \n", i)
    fmt.Println(i.Addtwo())
}
