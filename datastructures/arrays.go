package main

import "fmt"

func main(){

	var fruits [2]string = [2]string{"apples", "oranges"}
	fmt.Println(fruits)

	//short hand declaration
	grades := [3]int{10, 20, 30}
	fmt.Println(grades)

	names := [...]string{"Rachel", "Phoebe", "Monica"}
	fmt.Println(names)

	//length of the array
	fmt.Println(len(names))

	//access element at index 0 and 1 , as arrays starts with 0
	fmt.Println(names[0])
	fmt.Println(names[1])
}