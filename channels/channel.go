package main

import "fmt"

func main() {
	message := make(chan string)

	go func() { message <- "Ping" }()

	msg:= message;
	fmt.Println(string(<-msg));

}
