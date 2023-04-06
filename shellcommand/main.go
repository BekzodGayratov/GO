package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	command := exec.Command("ping","google.com")

	command.Stdout = os.Stdout;

	if err:= command.Run();err !=nil {
		fmt.Println("Could not run command: ",err)
	}
	

}
