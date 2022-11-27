package main

import "fmt"

func main() {
	// var a [11]int;
	// a[0]= 1;
	// a[1]= 2;
	// fmt.Println(len(a));

	// b := [5]int{1, 2, 3, 4, 5}

	// for i := 0; i < len(b); i++ {
	// 	fmt.Println(b[i])
	// }

	var myArray [1][1][1][1]string

	myArray[0][0][0][0] = "Bekzod"

	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}
}
