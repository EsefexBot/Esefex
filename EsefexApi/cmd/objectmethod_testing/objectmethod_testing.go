package main

import (
	"log"
)

type Routes struct {
	context string
}

func (s *Routes) MethodWithContext(arg string) {
	log.Printf("Context: %s, Arg: %s", s.context, arg)
}

func MethodWithoutContext(arg string) {
	log.Printf("Arg: %s", arg)
}

func RegisterHandler(handler func(string)) {
	handler("arg")
}

func main() {
	r := &Routes{"context"}
	RegisterHandler(r.MethodWithContext)

	RegisterHandler(MethodWithoutContext)
}
