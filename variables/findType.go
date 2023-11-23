package main

import "fmt"

func main() {
	var grades int = 42
	var message string = "hello world"
	var isCheck bool = true
	var amount float32 = 5466.32

	fmt.Printf("variable grades = %v is of type %T", grades, grades)
	fmt.Printf("variable message = %v is of type %T", message, message)
	fmt.Printf("variable isCheck = %v is of type %T", isCheck, isCheck)
	fmt.Printf("variable amount = %v is of type %T", amount, amount)
}
