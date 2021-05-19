package main

import (
	"fmt"
)

type Phone interface {
	call()
}

type MiPhone struct {
}

func (miPhone MiPhone) call() {
	fmt.Println("I am mi phone")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am i phone")
}

func main() {
	var phone Phone

	phone = new(MiPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}
