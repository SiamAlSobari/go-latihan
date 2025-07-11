package main

import (
	"fmt"
	"strconv"
)

func main() {
	age := 16
	str := strconv.Itoa(age) //convert int to string
	fmt.Println(str)


	
	num,err := strconv.Atoi(str) //convert string to int
	//error handling jika dia bukan angka string
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(num)
}