package main

import "fmt"

type Student struct {
	name string
	age  int
}

func (stu *Student) Hello() string {
	return fmt.Sprintf("name=>%v, age=>%v", stu.name, stu.age)
}
