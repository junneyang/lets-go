package main

import (
	"context"
	"grpc/user_proto"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8080"
)

type userServer struct {
	user_proto.UnimplementedUserServiceServer
	userDetails []*user_proto.UserDetail
}

func (userServer *userServer) GetUserDetail(context context.Context, user *user_proto.User) (*user_proto.UserDetail, error) {
	log.Printf("Received user=>%v", user)
	for _, userDetail := range userServer.userDetails {
		if user.GetName() == userDetail.GetName() {
			return userDetail, nil
		}
	}
	return &user_proto.UserDetail{}, nil
}

func (userServer *userServer) ListUserDetails(userGroup *user_proto.UserGroup, stream user_proto.UserService_ListUserDetailsServer) error {
	for _, user := range userGroup.GetUsers() {
		for _, userDetail := range userServer.userDetails {
			if userDetail.GetName() == user.GetName() {
				if err := stream.Send(userDetail); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (userServer *userServer) GetUserDetailGroup(stream user_proto.UserService_GetUserDetailGroupServer) error {
	userDetails := []*user_proto.UserDetail{}
	for {
		user, err := stream.Recv()
		log.Printf("Received user=>%v, err=%v", user, err)
		if err == io.EOF {
			userDetailGroup := &user_proto.UserDetailGroup{UserDetails: userDetails}
			return stream.SendAndClose(userDetailGroup)
		}
		if err != nil {
			return err
		}
		for _, userDetail := range userServer.userDetails {
			if userDetail.GetName() == user.GetName() {
				userDetails = append(userDetails, userDetail)
			}
		}
	}
}

func (userServer *userServer) UserChat(stream user_proto.UserService_UserChatServer) error {
	for {
		user, err := stream.Recv()
		log.Printf("Received user=>%v, err=%v", user, err)
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		user.Name = user.Name + "-S"
		if err := stream.Send(user); err != nil {
			return err
		}
	}
}

func (userServer *userServer) loadUserDetails() {
	userServer.userDetails = []*user_proto.UserDetail{
		{Name: "zhangsan", Address: "beijing", Phone: "18665817689"},
		{Name: "lisi", Address: "shenzhen", Phone: "18665817689"},
		{Name: "wangwu", Address: "shanghai", Phone: "18665817689"},
		{Name: "zhaogou", Address: "shanghai", Phone: "18665817689"},
		{Name: "liumao", Address: "shanghai", Phone: "18665817689"},
	}
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	userServer := &userServer{}
	userServer.loadUserDetails()
	user_proto.RegisterUserServiceServer(s, userServer)
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
