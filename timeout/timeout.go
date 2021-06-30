package main

import (
	"log"
	"time"
)

func test(ch chan bool) {
	time.Sleep(time.Second * 10)
	ch <- true
}

func wait(ch chan bool) {
	select {
	case <-ch:
		log.Println("success...")
		return
	case <-time.After(time.Second * 5):
		log.Println("timeout...")
		return
	}
}

func main() {
	ch := make(chan bool)
	log.Println("start...")
	go test(ch)
	wait(ch)
	log.Println("end...")
}
