package main

import "fmt"

func main() {
	fmt.Println("Hello, Golang")

	p := person{name: "John", age: 35}
	p.showName()
}
