package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Tick(time.Second)
	t2 := time.Tick(time.Second)
	var count int
	for {
		select {
		case <-t1:
			fmt.Println("咖啡色的羊驼")
			count++
			fmt.Println("count--->", count)
		case <-t2:
			fmt.Println(time.Now().Unix())
			count++
			fmt.Println("count--->", count)
		}
	}
}
