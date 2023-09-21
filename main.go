package main

import (
	"fmt"
	"seis-automation-box/cmd/controller"
	"seis-automation-box/cmd/middleware"
	"seis-automation-box/internal/envStaffs"
	optionalfeatures "seis-automation-box/internal/optionalFeatures"
)

func main() {
	fmt.Println("Current var:", middleware.CurrentDevice())
	essentialInfo := envStaffs.LoadEssentialInfoConfig("./configs/essentialInfo.json")
	fmt.Printf("%+v\n", essentialInfo)
	option := optionalfeatures.LoadOptionConfig("./configs/option.json")
	fmt.Printf("%+v\n", option)
	// steps
	// Validate configs.
	controller.ValidateConfig()
	// Only run the instRunner.
}
