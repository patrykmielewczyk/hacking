package main

import (
	"encoding/base64"
	"fmt"
	"os/exec"
)

func main() {
	encoded := "aWQ=" // base64 coded 'id'
	decoded, _ := base64.StdEncoding.DecodeString(encoded)

	cmd := exec.Command("/bin/sh", "-c", string(decoded))
	output, _ := cmd.Output()
	fmt.Println(string(output))
}
