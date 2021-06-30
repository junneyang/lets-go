package main

import "fmt"

func defer_demo() {
	// defer fmt.Println("Third")
	// fmt.Println("First")
	// fmt.Println("Second")
	if err := recover(); err != nil {
		fmt.Println(err)
	}
}
