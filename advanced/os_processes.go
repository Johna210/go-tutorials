package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main_task() {
	cmd := exec.Command("sleep", "5")

	// Start the command
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error starting command:", err)
		return
	}

	// Waiting
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Error waiting:", err)
		return
	}
	fmt.Println("Process is complete")
}

func simple_task_two() {
	cmd := exec.Command("grep", "foo")

	// Set input for the command
	cmd.Stdin = strings.NewReader("foo\nbar\nbaz\n")

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("Output:", string(output))
}

func simple_task() {
	cmd := exec.Command("echo", "Hello, World!")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("Output:", string(output))
}
