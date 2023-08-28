package main

import (
	"fmt"
	"runtime"
)

func main() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("You are running on macOS!")
	case "linux":
		fmt.Println("You are running on Linux!")
	case "windows":
		fmt.Println("You are running on Windows!")
	default:
		fmt.Printf("You are running on another OS: %s\n", os)
	}
}
