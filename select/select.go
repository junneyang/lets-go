package main

import (
	"fmt"
	"time"
)

func a1(ch chan bool) {
	time.Sleep(time.Second * 2)
	// ch <- true
	close(ch)
}

func a2(ch chan bool) {
	time.Sleep(time.Second * 5)
	ch <- true
}

func a3(ch chan bool) {
	time.Sleep(time.Second * 3)
	ch <- true
}

func main() {
	// timeout := make(chan bool)
	// ch := make(chan struct{})
	// go func() {
	// 	time.Sleep(time.Second * 5)
	// 	timeout <- true
	// }()
	// // for {
	// select {
	// case <-timeout:
	// 	log.Println("timeout...")
	// case <-ch:
	// 	log.Println("ch...")
	// 	// default:
	// 	// 	log.Println("default...")
	// }
	// // time.Sleep(time.Second * 1)
	// // }
	ch1 := make(chan bool, 0)
	ch2 := make(chan bool, 0)
	ch3 := make(chan bool, 0)
	go a1(ch1)
	go a2(ch2)
	go a3(ch3)
	select {
	case <-ch1:
		fmt.Printf("ch1 case...\n")
	case <-ch2:
		fmt.Printf("ch2 case...\n")
	case <-ch3:
		fmt.Printf("ch3 case...\n")
	}
}
