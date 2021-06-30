package student

import (
	"fmt"
	"log"
)

type StudentRequest struct {
	Name string
	Age  int
}

type StudentResponse struct {
	Message string
}

type StudentService interface {
	SayHello(req StudentRequest, resp *StudentResponse) error
}

type StudentServer struct{}

func (studentServer *StudentServer) SayHello(req StudentRequest, resp *StudentResponse) error {
	log.Printf("studentServer req=>%v\n", req)
	resp.Message = fmt.Sprintf("Hello, name=>%v, age=>%v", req.Name, req.Age)
	return nil
}
