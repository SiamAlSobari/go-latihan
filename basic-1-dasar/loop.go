package main

import (
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	wg.Add(2)

	go func(){
		loop_array()
		wg.Done()
	}()

	go func(){
		loop_sederhana()
		wg.Done()
	}()
	wg.Wait()
}

func loop_sederhana(){
	for i:= 1; i <= 1000; i++ {
		fmt.Println("ini perulangan ke", i)
	}
}

func loop_array(){
	arr:= []string{"siam", "budi", "andi", "joko"}
	for _, value := range arr{
		fmt.Println(value)
	}
}