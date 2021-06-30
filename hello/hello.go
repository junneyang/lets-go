package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	//    log.SetPrefix("greetings: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)

	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	//    message, err := greetings.Hello("Golang")
	//    message, err := greetings.Hello("")
	//    if err != nil {
	//        log.Fatal(err)
	//    }
	//    fmt.Println(message)

	//    names := []string{"Gladys", "Samantha", "Darrin"}
	names := []string{"Gladys", "Samantha", ""}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
