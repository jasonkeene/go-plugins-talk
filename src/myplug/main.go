package main

import (
	"C"
	"log"
)

func init() {
	log.Print("init for myplug was called")
}

func Foo() string {
	log.Print("Foo for myplug was called")
	return "bar"
}
