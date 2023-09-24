package controller

import (
	"fmt"
	"os"
	"seis-automation-box/internal/envStaffs"
)

func instEvnStaffs(system string) {
	// accourding os system, run essentail installation.
	switch system {
	case "darwin":
		// darwin is macOS.
		envStaffs.PlaceToolOnMac()
	case "linux":
		fmt.Println("Now is preparing Seis-automation-box for your Linux device...")
	case "windows":
		fmt.Println("Now is preparing Seis-automation-box for your Windows device...")
	default:
		fmt.Println("Error: Please check the middleware.")
		os.Exit(1)
	}
}

func instOptionFeatures(system string) {

}

func InstRunner(system string) {
	// Install essential staffs.
	instEvnStaffs(system)
	// accourding config, run optional installation.

}
