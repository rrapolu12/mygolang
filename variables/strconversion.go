package main

import (
	"fmt"
	"strconv"
)

func main(){
	var i int = 42
	var s string = strconv.Itoa(i) // convert int to string
	fmt.Printf("%q\n",s)

	i, err := strconv.Atoi(s)

	fmt.Printf("%v, %T \n", i, i)
	fmt.Printf("%v, %T \n", err, err)

	//error in conversion
	var ss string = "200abc"
	ii, err1 := strconv.Atoi(ss)

	fmt.Printf("%v, %T \n", ii, ii)
	fmt.Printf("%v, %T \n", err1, err1)
}