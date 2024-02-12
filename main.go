package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	for {
		fmt.Print("ccsh> ")
	
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "exit" {
			os.Exit(0)
		}

		cmd := exec.Command(input)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		cmd.Run()
	}
}