package main

import (
	"log"
	"reflect"
	"strings"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	t := reflect.TypeOf(wg)
	log.Println(t)
	log.Println(t.Elem().Name())
	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		argv := make([]string, 0, method.Type.NumIn())
		returns := make([]string, 0, method.Type.NumOut())
		for j := 0; j < method.Type.NumIn(); j++ {
			argv = append(argv, method.Type.In(j).Name())
		}
		for j := 0; j < method.Type.NumOut(); j++ {
			returns = append(returns, method.Type.Out(j).Name())
		}
		log.Printf("func (w *%s) %s(%s) %s",
			t.Elem().Name(),
			method.Name,
			strings.Join(argv, ","),
			strings.Join(returns, ","))
	}
}
