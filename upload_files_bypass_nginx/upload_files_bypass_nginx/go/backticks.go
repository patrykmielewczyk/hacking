package main

import (
	"fmt"
	"os/exec"
)

func main() {
	command := exec.Command("/bin/sh", "-c", "id")
	output, _ := command.Output()
	fmt.Println(string(output))
}
