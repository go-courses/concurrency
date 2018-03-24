package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			time.Sleep(10 * time.Second)
			fmt.Println(k)
			time.Sleep(10 * time.Second)
		}(i)
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, os.Kill)

	defer signal.Stop(signalChan)
	fmt.Println("main")
	go func() {
		p := <-signalChan
		fmt.Println(p)
		fmt.Println("Hey there")
		os.Exit(1)
	}()

	wg.Wait()

}
