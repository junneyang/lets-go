package main

type User struct {
	Name    string
	Age     int
	Address string
}

type Request struct {
	Name string
	Age  int
}

type Response struct {
	Message string
}
