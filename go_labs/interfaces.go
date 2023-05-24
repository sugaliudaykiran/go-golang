package main
import ("fmt"
        "math")
// Empty Interface..
type myinterface interface{}

//  "interface" is a custom type that is used to specify a set of one or more method signatures and the interface is abstract, so you are not allowed to create an instance of the interface. But you are allowed to create a variable of an interface type and this variable can be assigned with a concrete type value that has the methods the interface requires. Or in other words, the interface is a collection of methods as well as it is a custom type.
// it is necessary to implement all the methods declared in the interface for implementing an interface. The go language interfaces are implemented implicitly. And it does not contain any specific keyword to implement an interface just like other languages.

type shape interface{
    area() float64
}

type square struct{
    length float64
}

type circle struct{
    radius float64
}

// circle func..
func (c circle)area() float64{
    return math.Pi * c.radius * c.radius
}

// square func..
func (s square)area() float64{
    return s.length * s.length
}

func printInterface(s shape){ // Instead of defining which type like square, circle, mention the interface 
    fmt.Printf("area of %T is: %0.2f \n", s, s.area())    
}

func main(){
    shapes := []shape{ 
        square{length:15.2},
        circle{radius:12.2},
    }
    
    for _, v := range shapes{
        printInterface(v)
        fmt.Println("---")
    }
}
