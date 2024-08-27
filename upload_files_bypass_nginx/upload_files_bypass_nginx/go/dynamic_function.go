package main

import (
	"fmt"
	"os/exec"
)

func main() {
	funcName := "system" // we need to make function in GO
	cmd := exec.Command("/bin/sh", "-c", "id")
	output, _ := cmd.Output()
	fmt.Println(string(output))
}