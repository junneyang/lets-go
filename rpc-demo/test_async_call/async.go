package main

import (
	"log"
	"time"
)

// 简单示例，官方rpc框架有对call注册、传递的过程
type Call struct {
	Done chan *Call
}

var call = &Call{}

func (call *Call) done() {
	call.Done <- call
}

func Go(a int, b int, done chan *Call) *Call {
	if done == nil {
		done = make(chan *Call)
	}
	log.Println("Go start...")
	time.Sleep(time.Second * 5)
	call := &Call{Done: done}
	log.Println("Go end...")
	return call
}

func receive() {
	log.Println("receive start...")
	time.Sleep(time.Second * 10)
	log.Println("receive end...")
	call.done()
}

func main() {
	call = Go(10, 20, nil)
	go receive()
	<-call.Done
}
