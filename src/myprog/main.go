package main

import (
	"fmt"
	"log"
	"plugin"
)

func main() {
	p, err := plugin.Open("/go/bin/myplug.so")
	if err != nil {
		log.Panic(err)
	}
	s, err := p.Lookup("Foo")
	if err != nil {
		log.Panic(err)
	}

	f, ok := s.(func() string)
	if !ok {
		log.Panic("type assertion failed")
	}
	fmt.Printf("plugin call resulted in: %s\n", f())
}
