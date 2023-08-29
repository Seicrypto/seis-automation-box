package userDevice

import (
	"fmt"
	"runtime"
	"seis-automation-box/internal/errorHandle"
)

func CurrentDevice() string {
	var deviceOs string
	switch deviceOs = runtime.GOOS; deviceOs {
	case "darwin":
		fmt.Println("You are running on macOS!")
	case "linux":
		fmt.Println("You are running on Linux!")
		errorHandle.ExitWithError("Sorry, linux services are coming soon.")
	case "windows":
		fmt.Println("You are running on Windows!")
	default:
		fmt.Printf("You are running on another OS: %s\n", deviceOs)
		errorHandle.ExitWithError("Sorry, Seis-automation-box currently only support macOS, and Windows.")
	}
	return deviceOs
}
