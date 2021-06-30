package main

import (
	"fmt"
	"time"
)

func download_ch(ch chan string, url string) {
	defer func() { ch <- url }()
	fmt.Printf("start to download: %v\n", url)
	time.Sleep(time.Duration(5) * time.Second)
	fmt.Printf("end to download: %v\n", url)
	// ch <- url
}

func download_ch_once(ch chan string, url string) {
	defer func() { ch <- url }()
	fmt.Printf("start to download: %v\n", url)
	time.Sleep(time.Duration(5) * time.Second)
	fmt.Printf("end to download: %v\n", url)
	// ch <- url
}
