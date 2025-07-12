package main

import (
	"fmt"
)

func main(){
	arr:= make(map[string]string)
	arr["name"] = "siam"
	arr["age"] = "16"
	fmt.Println(arr)

	for _, value := range arr{
		fmt.Println(value)
	}

	arrName := []string{"siam", "budi", "andi", "joko"}

	for _, value := range arrName{
		fmt.Println(value)
	}
}