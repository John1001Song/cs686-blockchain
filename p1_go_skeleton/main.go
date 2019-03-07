package main

import (
	"fmt"
)

type person struct {
	name string
	age int
}

func main() {
	_ = person{
		name: "Joe",
		age:  20,
	}
	fmt.Println("hello world!")
}
