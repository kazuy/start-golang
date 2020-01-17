package main

import "fmt"

type person struct {
	name string
	age int
}

func (p *person) showName() {
	fmt.Printf("My name is %s.\n", p.name)
}
