package main

import "fmt"

func main(){
	// i:= 0;
	// for i<10{
	// 	fmt.Println(i);
	// 	i++
	// }

	// for j:= 0; j<=10;j++{
	// 	fmt.Println(j);
	// }

	// for{
	// 	fmt.Println("Loop");
	// 	break;
	// }

	for n:= 0; n<=5; n++{
		if(n%2==0){
			continue;
		}
		fmt.Println(n);
	}

}