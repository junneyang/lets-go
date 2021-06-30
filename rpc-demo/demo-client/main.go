package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"example.com/rpc-demo/codec"
)

func main() {
	conn, _ := net.Dial("tcp", ":8080")
	defer conn.Close()

	json.NewEncoder(conn).Encode(codec.DefaultOption)
	cc := codec.NewGobCodec(conn)
	header := &codec.Header{ServiceMethod: "Foo.Bar", Seq: 1001}
	cc.Write(header, fmt.Sprintf("rpc req %v", header.Seq))

	cc.ReadHeader(header)
	var resp string
	cc.ReadBody(&resp)
	log.Printf("header=>%v, resp=%v", header, resp)
}
