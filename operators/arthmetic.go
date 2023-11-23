package main

import "fmt"

func main(){

	var a,b string = "foo", "bar"
	fmt.Println(a + b)

	//try substraction, you cannot define on string
	//fmt.Println(a - b)

	//try on int
	var i, j int = 20 , 10
	fmt.Println( i - j)

	//multiply
	fmt.Println(i * j)

	//division#
	fmt.Println(i/j)

	//modulus
	fmt.Println(i % j)
	//increment
	i++
	fmt.Println(i)
	//decrement
	i--
	fmt.Println(i)
}