package main
import ("fmt")

func main(){
    fmt.Println("Hello world..!")
    // variable 
    var st_name string = "Kiran";
    var st_age int = 22; 
    var st_phone = "800813971";  // inferred type of declaration..
    st_gender := "male";
	// var st_class, st_section int = 1,2;  // declaring the similar type of multiple variables..
    st_class, st_section := 1,"B";  // declaring the multiple variable without using "type and var"
    
    fmt.Println(st_name,st_age,st_phone,st_gender,st_class,st_section);

	// constant variable
	const PI float32 = 3.14;
	fmt.Println(PI);

	const(
        A = 1;
        B = 2;
        // A = 10;
    )
    fmt.Println(A,B);
}