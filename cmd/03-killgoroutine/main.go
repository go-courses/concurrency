package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	stopChan := make(chan bool, 2)
	wg.Add(1)
	stopChan <- true
	go listen(stopChan)

	doStuff(stopChan)
	wg.Wait()
}

func doStuff(c chan bool) {
	time.Sleep(5 * time.Second)
	fmt.Println("exit")
	close(c)
}

func listen(c chan bool) {
	for range time.Tick(time.Second) {
		select {
		case b := <-c:
			if b {
				fmt.Println("something sent")
				continue
			}
			wg.Done()
			return
		default:
			fmt.Println("listening...")
		}
	}
}
