package main

import "fmt"

func main() {

	arr := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	
	
	//element at 0,0
	fmt.Println(arr[0][0])
	
	//element at 0,1
	fmt.Println(arr[0][1])

	//element at 3,3
	fmt.Println(arr[2][1])
}
