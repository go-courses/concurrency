package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	start()
	wg.Wait()
}
func start() {
	wg.Add(1)
	go listenRedis()
	wg.Add(1)
	go listenDatabase()
	wg.Add(1)
	go listenHTTP()
}
func listenRedis() {
	for range time.Tick(time.Second) {
		fmt.Println("Hey from redis")
	}
	wg.Done()
}
func listenDatabase() {
	for range time.Tick(time.Second) {
		fmt.Println("Hey from database")
	}
	wg.Done()
}
func listenHTTP() {
	for range time.Tick(time.Second) {
		fmt.Println("Hey from HTTP")
	}
	wg.Done()
}
