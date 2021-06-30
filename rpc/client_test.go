package main

import (
	"crypto/tls"
	"log"
	"net/rpc"
	"testing"
)

func TestClient(t *testing.T) {
	config := &tls.Config{
		InsecureSkipVerify: true,
	}
	conn, _ := tls.Dial("tcp", "localhost:8080", config)
	defer conn.Close()
	client := rpc.NewClient(conn)
	// client, _ := rpc.DialHTTP("tcp", "localhost:8080")
	req := Request{Name: "zhangsan", Age: 18}

	// resp := &Response{}
	// if err := client.Call("UserServer.SayHello", req, resp); err != nil {
	// 	log.Println(err)
	// }
	// log.Printf("req=>%v, resp=%v\n", req, resp)

	// resp = &Response{}
	// async := client.Go("UserServer.SayHello", req, resp, nil)
	// log.Printf("req=>%v, resp=%v\n", req, resp)
	// <-async.Done
	// log.Printf("req=>%v, resp=%v\n", req, resp)

	resp := &Response{}
	asyncStudent := client.Go("StudentServer.SayHello", req, resp, nil)
	log.Printf("req=>%v, resp=%v\n", req, resp)
	<-asyncStudent.Done
	log.Printf("req=>%v, resp=%v\n", req, resp)
}
