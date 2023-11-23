package main

import "fmt"

func main(){
	for i := 1; i<=3; i++{
		fmt.Println("Hello World!")
	}

	for ii := 1; ii <= 3; ii++{
		fmt.Println(ii * ii)
	}

	//use a break statement
	for i := 1; i<=5; i++{
		if i == 3 {
			break
		}
		fmt.Println("Hello World!")
	}

	fmt.Println("=======================")
		//use a break statement
		for i := 1; i<=5; i++{
			if i == 3 {
				continue
			}
			fmt.Println("Hello World!")
		}

}