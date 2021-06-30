package main

import (
	"log"
	"sync"
	"time"
)

func test(wg *sync.WaitGroup) {
	log.Println("start sleep test...")
	time.Sleep(time.Second * 20)
	log.Println("end sleep test...")
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go test(wg)
	log.Println("start sleep...")
	time.Sleep(time.Second * 5)
	log.Println("end sleep...")
	wg.Wait()
}
