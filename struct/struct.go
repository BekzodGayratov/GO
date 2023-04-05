package main

import "fmt"

func main() {
	p := Tesla{model: "X10", number: "801"}
	fmt.Println(p.model)

}

type Tesla struct {
	model  string
	number string

	
}
