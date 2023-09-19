package main

import (
	"fmt"
	"seis-automation-box/cmd/middleware"
	"seis-automation-box/internal/envStaffs"
	optionalfeatures "seis-automation-box/internal/optionalFeatures"
)

func main() {
	fmt.Println("Current var:", middleware.CurrentDevice())
	essentialInfo := envStaffs.LoadEssentialInfoConfig()
	fmt.Printf("%+v\n", essentialInfo)
	option := optionalfeatures.LoadOptionConfig()
	fmt.Printf("%+v\n", option)
	// steps
	// Only run the instRunner.
}
