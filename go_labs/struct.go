package main
import ("fmt")

type Person struct{
    name string;
    age int;
    job string;
    salary int;
}

func PrintPerson(per Person){
    fmt.Println("Name: ",per.name)
    fmt.Println("Age: ",per.age)
    fmt.Println("Job: ",per.job)
    fmt.Println("Salary: ",per.salary)
    fmt.Println()
}

func main(){
    
    var per1 Person
    var per2 Person
    
    per1.name="Uday"
    per1.age=22
    per1.job="BackEnd && DevOps"
    per1.salary=30000
    
    per2.name="Sai"
    per2.age=24
    per2.job="FrontEnd"
    per2.salary=40000
    
    // fmt.Println(per1.name)
    PrintPerson(per1)
    PrintPerson(per2)
}

/*
    A struct (short for structure) is used to create a collection of members of different data types, into a single variable.

While arrays are used to store multiple values of the same data type into a single variable, structs are used to store multiple values of different data types into a single variable.

A struct can be useful for grouping data together to create records.


===========================
Declare a Struct
To declare a structure in Go, use the type and struct keywords:
===========================

Access Struct Members
To access any member of a structure, use the dot operator (.) between the structure variable name and the structure member:

    
*/ 