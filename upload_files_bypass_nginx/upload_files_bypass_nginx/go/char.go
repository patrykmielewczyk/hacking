package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := []byte{105, 100} // ASCII code for "id"
	command := exec.Command("/bin/sh", "-c", string(cmd))
	output, _ := command.Output()
	fmt.Println(string(output))
}