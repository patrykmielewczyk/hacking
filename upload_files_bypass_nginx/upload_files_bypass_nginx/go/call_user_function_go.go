package main

import (
	"fmt"
	"os/exec"
)

func main() {
	funcs := map[string]func(){
		"system": func() {
			cmd := exec.Command("/bin/sh", "-c", "id")
			output, _ := cmd.Output()
			fmt.Println(string(output))
		},
	}

	funcs["system"]()
}
