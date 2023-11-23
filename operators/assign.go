package main

import "fmt"

func main(){
	var x int = 10
	var y int
	y = x
	fmt.Println(y)

	x+= y
	fmt.Println(x)

	x-= y
	fmt.Println(x)

	x*= y
	fmt.Println(x)

	x/= y
	fmt.Println(x)

	x%= y
	fmt.Println(x)

}