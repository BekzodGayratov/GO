package main

import "fmt"

func main() {
	
	fmt.Println(Person{name: "Bekzod",age: 19})
}

type Person struct {
	name string
	age  int
}

func newPerson(name string) *Person {

	p := Person{name: name}
	p.age = 42
	return &p
}
