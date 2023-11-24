package main

import (
	"fmt"
)

func main(){
	
	var codes map[string]string
	fmt.Println(codes)
	
	languages := map[string]string{"en":"English","fr":"French","hi":"Hindi"}
	fmt.Println(languages)
	
	value,found := languages["en"]
	fmt.Println(value,found)
	//update
	languages["en"] = "English Language"
	fmt.Println(languages)

	//delete
	delete(languages, "en")
	fmt.Println(languages)
	//add
	languages["en"] = "English Language"
	fmt.Println(languages)

	//loop
	for key,value := range languages{
		fmt.Println(key,"=>",value)
	}
	//truncate
	languages = make(map[string]string)
	fmt.Println(languages)
}