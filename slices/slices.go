package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s[0])

	s = append(s, "Bekzod")

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)

	l := s[2:5]
	fmt.Println(l)

}
