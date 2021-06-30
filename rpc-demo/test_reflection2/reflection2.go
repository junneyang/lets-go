package main

import (
	"fmt"
	"log"
	"reflect"
)

type Request struct {
	Name string
	Age  int
}

type Response struct {
	Message string
}

type HelloService interface {
	SayHello(req Request, resp *Response) error
}

type UserServer struct{}

func (userServer *UserServer) SayHello(req Request, resp *Response) error {
	log.Printf("userServer req=>%v\n", req)
	resp.Message = fmt.Sprintf("Hello, name=>%v, age=>%v", req.Name, req.Age)
	return nil
}

type MethodInfo struct {
	n string
	m reflect.Method
	t reflect.Type
	v reflect.Value
}

var serviceMap = make(map[string]*MethodInfo)

func Register(service interface{}) {
	t := reflect.TypeOf(service)
	serviceName := t.Elem().Name()
	log.Println(serviceName)
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		methodName := method.Name
		log.Println(methodName)
		name := serviceName + "." + methodName
		methodInfo := &MethodInfo{n: methodName, m: method, t: reflect.TypeOf(service), v: reflect.ValueOf(service)}
		serviceMap[name] = methodInfo
	}
	log.Println(serviceMap)
}

func main() {
	Register(new(UserServer))
	service := serviceMap["UserServer.SayHello"]
	f := service.m.Func
	request := Request{Name: "zhangsan", Age: 18}
	response := &Response{}
	f.Call([]reflect.Value{service.v, reflect.ValueOf(request), reflect.ValueOf(response)})
	log.Println(response)
}
