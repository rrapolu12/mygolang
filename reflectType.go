package main

import (
	"fmt"
	"reflect"
)

func main() {
	var grades int = 42
	var message string = "hello world"

	fmt.Printf("variable grades = %v is of type %T \n", grades, reflect.TypeOf(grades))
	fmt.Printf("variable message = %v is of type %T \n ", message, reflect.TypeOf(message))
}
