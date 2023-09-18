package middleware

import (
	"fmt"
	"os"
	"runtime"
)

func CurrentDevice() string {
	var userDevice string
	switch userDevice = runtime.GOOS; userDevice {
	case "darwin":
		fmt.Println("Now is preparing Seis-automation-box for your Mac device...")
	case "linux":
		fmt.Println("Now is preparing Seis-automation-box for your Linux device...")
	case "windows":
		fmt.Println("Now is preparing Seis-automation-box for your Windows device...")
	default:
		fmt.Println("Error: Sorry, Seis-automation-box currently only support macOS, and Windows.")
		os.Exit(1)
	}
	return userDevice
}
