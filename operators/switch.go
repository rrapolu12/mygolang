package main

import "fmt"

func main() {

	var i int = 0

	switch i {
	case 0:
		fmt.Println("Number is 0")
	case 100, 200:
		fmt.Println("Number is either 100 or 200")
	default:
		fmt.Println("i is neither 0,100,200")
	}

	var ii int = 10

	switch ii {
	case -5:
		fmt.Println("-5")
	case 10:
		fmt.Println("10")
		fallthrough
	case 0:
		fmt.Println("0")
		fallthrough
	default:
		fmt.Println("default")

	}

	fmt.Println("===========================")
	var a, b int = 10, 20
	switch {
	case a+b == 30:
		fmt.Println("equal to 30")
	case a+b <= 30:
		fmt.Println("less than or equal to 30")
	default:
		fmt.Println("greater than 30")
	}

}
