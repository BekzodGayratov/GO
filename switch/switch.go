package main

import "fmt"

func main(){

	i:= 2;

	switch i{
	case 1:
		fmt.Println("bir");
		break;
	case 2:
		fmt.Println("ikki");
		break;
	case 3: 
		fmt.Println("uch");
		break;
	default:
		fmt.Println("noma'lum");
	}
}