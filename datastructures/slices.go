package main

import "fmt"

func main(){
	
	

	arr := [10]int{10,20,30,40,50,60,70,80,90,100}

	intslice := arr[1:8]
	fmt.Println(intslice)

	//another way of initializing slice

	slice := make([]int,5,10)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))



	fmt.Println("====================================NUMBERS=============================================")
	//append
	numbers := [4]int{1,2,3,4}
	fmt.Println(numbers)

	n_slice := numbers[1:3]
	fmt.Println(n_slice)
	fmt.Println(len(n_slice))
	fmt.Println(cap(n_slice))

	fmt.Println("=================update========================")
	fmt.Println(numbers)
	n_slice[1] = 100
	fmt.Println(numbers)
	fmt.Println(n_slice)

	fmt.Println("=================append========================")
	n_slice = append(n_slice, 5,6,7,8)
	fmt.Println(n_slice)
	fmt.Println(len(n_slice))
	fmt.Println(cap(n_slice))

	

	fmt.Println()
}