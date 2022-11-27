package main

import (
	"fmt"
)

func main() {
	fmt.Println(orttir(1, 2, 3));
	fmt.Println(kamaytir(5,3));
}

func orttir(a, b, c int) int {
	return a + b + c;
}

func kamaytir (a,b int)int{
	return a-b;
}
