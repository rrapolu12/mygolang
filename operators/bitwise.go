package main

import "fmt"

func main(){

	//bitwise AND
	var x, y int = 12 , 25
	var z = x & y
	fmt.Println(z)

	//bitwise OR

	z = x | y
	fmt.Println(z)

	//bitwise XOR(^)
	z = x ^ y
	fmt.Println(z)

	//bitwise left shift
	z = x << 1
	fmt.Println(z)

	//bitwise right shift
	z = x >> 1
	fmt.Println(z)
	

}