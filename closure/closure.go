package main

import "fmt"

func main() {
	nextInt := sum();
	fmt.Println(nextInt());
	fmt.Println(nextInt());
	fmt.Println(nextInt());
	
}

func sum() func() int {
	i:= 0;
	return func() int {
		i++
		return i
	}
}
