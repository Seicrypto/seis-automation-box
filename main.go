package main

import (
	"seis-automation-box/cmd/controller"
	"seis-automation-box/cmd/middleware"
)

func main() {
	// Test
	// fmt.Println("Current var:", middleware.CurrentDevice())
	// essentialInfo := envStaffs.LoadEssentialInfoConfig("./configs/essentialInfo.json")
	// fmt.Printf("%+v\n", essentialInfo)
	// option := optionalfeatures.LoadOptionConfig("./configs/option.json")
	// fmt.Printf("%+v\n", option)

	// steps
	// Validate configs,
	// Be care! go run command would execute from different path between built file.
	controller.ValidateConfig()
	// Only run the instRunner.
	controller.InstRunner(middleware.CurrentDevice())
}
