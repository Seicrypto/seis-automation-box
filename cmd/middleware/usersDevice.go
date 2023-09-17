package middleware

import (
	"fmt"
	"runtime"
	"seis-automation-box/internal/errorHandle"
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
		errorHandle.ExitWithError("Sorry, Seis-automation-box currently only support macOS, and Windows.")
	}
	return userDevice
}
