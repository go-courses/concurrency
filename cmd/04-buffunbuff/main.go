package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 2)
	c <- 0
	c <- 1
	go func(c chan int) {
		for range time.Tick(time.Second) {
			fmt.Println(<-c)
		}

	}(c)
	c <- 2
	c <- 4
	time.Sleep(5 * time.Second)
}
