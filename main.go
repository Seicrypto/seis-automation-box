package main

import (
	"fmt"
	"seis-automation-box/cmd/controller"
	"seis-automation-box/cmd/middleware"
)

func main() {
	fmt.Println("Current var:", middleware.CurrentDevice())
	essentialInfo := controller.LoadEssentialInfoConfig()
	fmt.Printf("%+v\n", essentialInfo)
	option := controller.LoadOptionConfig()
	fmt.Printf("%+v\n", option)
	// steps
	// Only run the instRunner.
}
