package main

import "fmt"



func main(){
	a,b,c := vals();
	
	fmt.Println(a,b,c);
}


func vals()(int,int,int){
	return 3,5,6;
}

