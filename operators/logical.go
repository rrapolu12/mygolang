package main

import "fmt"

func main(){
	var x int = 10
	fmt.Println((x<100) && (x < 200))
	fmt.Println((x<300) && (x < 0))

	// Logical OR ||

	fmt.Println((x<100) || (x < 200))
	fmt.Println((x<300) || (x < 0))

	fmt.Println(!(x<100))
	fmt.Println(!(x<0))
}