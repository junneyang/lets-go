package main

import (
	"context"
	"grpc/user_proto"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:8080"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := user_proto.NewUserServiceClient(conn)

	// 一次性调用
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3600)
	defer cancel()
	user := &user_proto.User{Name: "zhangsan"}
	r, err := c.GetUserDetail(ctx, user)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r)

	// Server-Streaming
	users := []*user_proto.User{
		{Name: "zhangsan"},
		{Name: "lisi"},
		{Name: "wangwu"},
		{Name: "zhaogou"},
	}
	userGroup := &user_proto.UserGroup{Users: users}
	stream, err := c.ListUserDetails(ctx, userGroup)
	if err != nil {
		log.Fatalf("%v.ListUserDetails(_) = _, %v", c, err)
	}
	for {
		userDetail, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListUserDetails(_) = _, %v", c, err)
		}
		log.Printf("ListUserDetails=>%v", userDetail)
	}

	// Client Streaming
	c_stream, err := c.GetUserDetailGroup(ctx)
	if err != nil {
		log.Fatalf("%v.GetUserDetailGroup(_) = _, %v", c, err)
	}
	for _, user := range users {
		if err := c_stream.Send(user); err != nil {
			log.Fatalf("%v.Send(%v) = %v", c_stream, user, err)
		}
	}
	userDetailGroup, err := c_stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.Recv() got error %v, want %v", c_stream, err, nil)
	}
	log.Printf("GetUserDetailGroup: %v", userDetailGroup)

	// 双向通信
	b_stream, err := c.UserChat(ctx)
	if err != nil {
		log.Fatalf("%v.UserChat(_) = _, %v", c, err)
	}
	wait := make(chan struct{})
	go func() {
		for {
			user, err := b_stream.Recv()
			log.Printf("Received user=>%v, err=%v", user, err)
			if err == io.EOF {
				close(wait)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a user : %v", err)
			}
		}
	}()
	for _, user := range users {
		user.Name = user.Name + "-C"
		if err := b_stream.Send(user); err != nil {
			log.Fatalf("Failed to send a user: %v", err)
		}
		time.Sleep(time.Second * 10)
	}
	b_stream.CloseSend()
	<-wait
}
