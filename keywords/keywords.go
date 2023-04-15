package main

import (
	"fmt"
	"os"
)

func main() {
	//GO da 25 ta keyword mavjud.
	// break, case, chan, const, continue, default, defer, else, fallthrough, for, func, go, goto,
	// if, import, interface, map, package, range, return, select, struct, switch, type, var.

	// 1. break

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
			break
		}
	}

	// 2. case
	var age int16 = 18
	switch age {
	case 18:
		fmt.Println("This age is minimal to work")
		break
	}

	// 3. chan
	message := make(chan string)

	go func() { message <- "Ping" }()

	msg := message
	fmt.Println(string(<-msg))

	// 4. const

	var name string = "Bekzod"
	fmt.Println(name)

	// 5. continue

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		} else {
			continue
		}
	}

	// 6. default

	age2 := 18

	switch age2 {
	case 15:
		fmt.Println("The minimal age to work")
		break
	default:
		fmt.Println("The age is unknown")
		break
	}

	// 7. defer
	f := createFile("/text.txt")
	defer closeFile(f)

	// 

}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
