package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// like assert("system('id');")
	cmd := exec.Command("/bin/sh", "-c", "id")
	output, _ := cmd.Output()
	fmt.Println(string(output))
}
