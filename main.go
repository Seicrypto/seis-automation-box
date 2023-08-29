package main

import (
	"fmt"
	"seis-automation-box/internal/userDevice"
)

func main() {
	fmt.Println("Current var:", userDevice.CurrentDevice())
}
