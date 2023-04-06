package main

import "fmt"

func main() {
	// int
	var x uint8 = 225
	fmt.Printf("The result is:%d\n", x)
	var y uint16 = 170
	var X uint16 = 20
	fmt.Printf("The result is: %d\n", y-X)

	// float
	var son1 float32 = 20.0
	var son2 float32 = 5
	fmt.Printf("The result is: %f\n", son1/son2)

	var son3 float64 = 50
	var son4 float64 = 5
	fmt.Printf("The result is: %f\n", son3-son4)

	// complex
	var a complex128 = complex(6, 2)
	var b complex64 = complex(9, 2)
	fmt.Println(a)
	fmt.Println(b)

	// booleans
	var isMarried bool = true
	fmt.Printf("The result is: %t", isMarried)

	// strings
	var name string = "Bekzod"
	fmt.Printf("The name is: %s\n", name)
	fmt.Printf("The length of name is: %d\n", len(name))

}
