package main

import "fmt"

func main(){

	var fruits [5]string = [5]string{"apples", "oranges","grapes","mango","papaya"}
	fmt.Println(fruits)

	//short hand declaration
	grades := [3]int{10, 20, 30}
	fmt.Println(grades)

	names := [...]string{"Rachel", "Phoebe", "Monica"}
	fmt.Println(names)

	//length of the array
	fmt.Println(len(names))

	//access element at index 0 and 1 , as arrays starts with 0
	fmt.Println(fruits[0])
	fmt.Println(fruits[1])

	//access element 4
	fmt.Println(fruits[3])

	//update element of names at index 0
	fmt.Println("===================================")
	fmt.Println(names)
	names[0] = "Sham"
	fmt.Println(names)

	//loop through the array

	fmt.Println("================LOOP===================")
	for i :=0; i<len(fruits); i++ {
		fmt.Println(fruits[i])
	}

	fmt.Println("================RANGE===================")

	for index, elemenet := range fruits {
		fmt.Println(index, "=>", elemenet)
	}

}