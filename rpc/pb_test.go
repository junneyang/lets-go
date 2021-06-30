package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"rpc/student"
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestPb(t *testing.T) {
	filename := "student.pb"

	// PB序列化
	// stu := &student.Student{
	// 	Name:   "zhnagsan",
	// 	Age:    18,
	// 	Sex:    student.Student_FEMALE,
	// 	Scores: []int32{98, 100, 99},
	// }
	// out, _ := proto.Marshal(stu)
	// filename := "student.pb"
	// ioutil.WriteFile(filename, out, 0644)

	// PB序列化
	// stu := &student.Student{
	// 	Name: "zhnagsan",
	// 	// Age:    18,
	// 	Sex:    student.Student_FEMALE,
	// 	Scores: []int32{98, 100, 99},
	// }
	// out, _ := proto.Marshal(stu)
	// filename := "student.pb"
	// ioutil.WriteFile(filename, out, 0644)

	// PB反序列化
	stustu := &student.Student{}
	in, _ := ioutil.ReadFile(filename)
	proto.Unmarshal(in, stustu)
	fmt.Println(stustu)
	fmt.Println(stustu.GetScores())
	fmt.Println(stustu.GetSex())
	log.Println(stustu)
}
