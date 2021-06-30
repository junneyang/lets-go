package main

import "fmt"

type HelloService interface {
	GetName() string
	SayHello() string
}

type Teacher struct {
	name   string
	age    int
	course string
}

type Docter struct {
	name       string
	age        int
	department string
}

func (teacher *Teacher) GetName() string {
	return teacher.name
}

func (teacher *Teacher) SayHello() string {
	return fmt.Sprintf("name=>%v, age=%v, course=>%v", teacher.name, teacher.age, teacher.course)
}

func (docter *Docter) GetName() string {
	return docter.name
}

func (docter *Docter) SayHello() string {
	return fmt.Sprintf("name=>%v, age=%v, department=>%v", docter.name, docter.age, docter.department)
}
