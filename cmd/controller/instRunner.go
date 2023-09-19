package controller

import (
	"fmt"
	"os"
	"seis-automation-box/internal/envStaffs"
)

func InstRunner(system string) {
	// accourding os system, run essentail installation.
	switch system {
	case "darwin":
		// darwin is macOS.
		fmt.Println("Now is preparing Seis-automation-box for your Mac device...")
		envStaffs.PlaceToolOnMac()
	case "linux":
		fmt.Println("Now is preparing Seis-automation-box for your Linux device...")
	case "windows":
		fmt.Println("Now is preparing Seis-automation-box for your Windows device...")
	default:
		fmt.Println("Error: Please check the middleware.")
		os.Exit(1)
	}
	// accourding config, run optional installation.
	switch system {
	case "darwin":
		// darwin is macOS.
		fmt.Println("Now is preparing option setting Mac device...")
	case "linux":
		fmt.Println("Now is preparing option setting Linux device...")
	case "windows":
		fmt.Println("Now is preparing option setting Windows device...")
	default:
		fmt.Println("Error: Please check the middleware.")
		os.Exit(1)
	}
	fmt.Println("Dev test finished.")
}
