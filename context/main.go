package main

import (
	"context"
	"log"
	"runtime"
	"time"
)

func demo(ctx context.Context, a string, b string) string {
	options := ctx.Value(Options{}).(*Options)
	// log.Printf("name=>%v, age=>%v", ctx.Value("name"), ctx.Value("age"))
	log.Printf("name=>%v, age=>%v", options.Name, options.Age)
	c := make(chan string, 1)
	go func() {
		log.Println("start sleep...")
		time.Sleep(time.Second * 10)
		log.Println("end sleep...")
		// goroutine stuck here
		c <- (a + " " + b)
		log.Println("end sleep......")
	}()

	select {
	case <-ctx.Done():
		return ctx.Err().Error()
	case res := <-c:
		return res
	}
}

type Options struct {
	Name string
	Age  int
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	// ctx = context.WithValue(ctx, "name", "HelloWorld")
	// ctx = context.WithValue(ctx, "age", 18)
	ctx = context.WithValue(ctx, Options{}, &Options{Name: "HelloWorld", Age: 18})
	defer cancel()

	log.Println("start to call...")
	res := demo(ctx, "Hello", "World")
	log.Printf("res=>%v\n", res)

	time.Sleep(time.Second * 10)
	log.Println("the number of goroutines: ", runtime.NumGoroutine())
}
