package main
import ("fmt")

func main(){
    // Arithmetic Operator :=
    x,y := 10,20
    
    fmt.Println(x+y, y-x, x*2, y/2, y%2)
    y++;x--
    fmt.Println(y, x)
    
    // Assignment Operator :=
    var a,b,c = 21, 11, 2
    // a+=b, a-=10, b*=2, b/=2, a%=10, c|=1, c&=2, c^=3, c>>=1, c<<=2
    fmt.Println(a, b, c);
    
    // Comparison Operator :=
    x1, x2, x3 := 11, 24, 10
    fmt.Println(x2>x1, x3<x1, x1!=x2, x2==x2, x3>=x1, x2<=x1)
    
    // Logical Operator :=
    y1, y2, y3 := true, true, false
    fmt.Println(y1&&y2, y2||y3, !y3)
    
    // Bitwise Operator :=
    z1, z2, z3 := 2, 5, 6
    fmt.Printf("binary format:\n z1&0 : %b, z2|1 : %b, z3^2 : %b, z2<<1 : %b, z1>>1 : %b \n\n",z1&0, z2|1, z3^2, z2<<1, z1>>1)
    fmt.Printf("binary format:\n z1&0 : %v, z2|1 : %v, z3^2 : %v, z2<<1 : %v, z1>>1 : %v \n\n",z1&0, z2|1, z3^2, z2<<1, z1>>1)
    
}