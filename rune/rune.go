package main

import (
	"fmt"
	"unicode/utf8"
)

func main(){
	const s = "salom"

	fmt.Println(len(s));

	for i:=0;i<len(s);i++{
		fmt.Println(s[i])
	}
	fmt.Println()

	fmt.Println("Rune count",utf8.RuneCountInString(s));
}