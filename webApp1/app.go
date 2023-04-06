package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/",sayHello);
	PORT := "3001";
	HOST := "localhost";	
	fmt.Printf("Your server running on: http://%s:%s",HOST,PORT);
	http.ListenAndServe(HOST+":"+PORT,nil);
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Golang :)")
}
