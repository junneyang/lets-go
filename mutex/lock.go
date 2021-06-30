package main

import (
	"log"
	"sync"
)

func add(count *int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	for i := 0; i < 10000; i++ {
		mutex.Lock()
		*count++
		mutex.Unlock()
	}
	wg.Done()
}

func main() {
	mutex := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	count := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go add(&count, wg, mutex)
	}
	wg.Wait()
	log.Printf("count=>%v\n", count)
}
