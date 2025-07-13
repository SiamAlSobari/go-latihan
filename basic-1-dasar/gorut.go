package main

import "fmt"

func kirim_data(c chan string) {
	for i := 0; i < 10; i++ {
		c <- "halo"
	}
	close(c)
}

func main() {
	c := make(chan string, 2)
	go kirim_data(c)
	c <- "hiiiiiiiiiiiii"
	fmt.Println(<-c)
	for pesan := range c {
		fmt.Println(pesan)
	}
}
