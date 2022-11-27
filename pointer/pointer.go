package main

import "fmt"

func main(){
	i:= 1;

	fmt.Println("Initial value",i);
	zeroval(i);
	fmt.Println("Zero value",i);
	zeroptr(&i);
	fmt.Println("Zeroptr",i);
	

}

func zeroval(ival int){
	ival = 0;
}

func zeroptr(iptr *int){
	*iptr = 0;
}
