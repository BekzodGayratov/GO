package main

import "fmt"

func main() {

	var person1 Person
	person1 = Person{name: "Hello", age: "18"}

	fmt.Println(person1.age)
}

type Person struct {
	name string
	age  string
}
