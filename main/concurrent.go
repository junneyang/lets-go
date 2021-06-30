package main

import (
	"fmt"
	"sync"
	"time"
)

func download(wg *sync.WaitGroup, url string) {
	defer wg.Done()
	fmt.Printf("start to download: %v\n", url)
	time.Sleep(time.Duration(5) * time.Second)
	fmt.Printf("end to download: %v\n", url)
	// wg.Done()
}
