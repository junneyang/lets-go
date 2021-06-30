package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/rpc"
	"rpc/student"
)

type HelloService interface {
	SayHello(req Request, resp *Response) error
}

type UserServer struct{}

func (userServer *UserServer) SayHello(req Request, resp *Response) error {
	log.Printf("userServer req=>%v\n", req)
	resp.Message = fmt.Sprintf("Hello, name=>%v, age=>%v", req.Name, req.Age)
	return nil
}

func main() {
	rpc.Register(new(UserServer))
	rpc.Register(new(student.StudentServer))
	cert, _ := tls.LoadX509KeyPair("server.crt", "server.key")
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	listener, _ := tls.Listen("tcp", ":8080", config)
	// rpc.HandleHTTP()
	log.Printf("Serving RPC server on port %d", 8080)
	// http.ListenAndServe(":8080", nil)
	for {
		conn, _ := listener.Accept()
		// defer conn.Close()
		go rpc.ServeConn(conn)
	}
}
