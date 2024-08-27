package main

import (
	"fmt"
	"os/exec"
)

func main() {
	part1 := "i"
	part2 := "d"
	command := exec.Command("/bin/sh", "-c", part1+part2)
	output, _ := command.Output()
	fmt.Println(string(output))
}
