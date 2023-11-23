package main

import "fmt"

func main(){
	var fruit string = "grapes"

	if fruit == "grapes"{
		fmt.Println("Fruit is Grapes")
	}else{
		fmt.Println("Fruit is not grapes")
	}

	var anotherFruit string = "apple"
	if (anotherFruit == "apple"){
		fmt.Println("I love apples")
	}else if anotherFruit == "Orange" {
		fmt.Println("i love oranges")
	}else if anotherFruit == "grapes" {
		fmt.Println("I love grapes")
	}else{
		fmt.Println("No fruit")
	}
}