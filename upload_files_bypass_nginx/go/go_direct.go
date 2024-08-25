package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// short php tag analogy
	cmd := exec.Command("/bin/sh", "-c", "id")
	output, _ := cmd.Output()
	fmt.Println(string(output))
}
