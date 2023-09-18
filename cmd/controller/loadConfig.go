package controller

import (
	"encoding/json"
	"fmt"
	"os"
)

type UserConfig map[string]string

type FeaturesConfig map[string]bool

func LoadConfigJSON() (UserConfig, FeaturesConfig) {
	var essential UserConfig
	var option FeaturesConfig

	essentialBytes, err := os.ReadFile("../../configs/essentialInfo.json")
	if err != nil {
		fmt.Println("Error reading essentialInfo.json:", err)
		os.Exit(1)
	}
	err = json.Unmarshal(essentialBytes, &essential)
	if err != nil {
		fmt.Printf("Error unmarshalling essentialInfo.json: %v\n", err)
		os.Exit(1)
	}

	optionBytes, err := os.ReadFile("../../configs/optionInst.json")
	if err != nil {
		fmt.Println("Error reading optionInst.json:", err)
		os.Exit(1)
	}
	err = json.Unmarshal(optionBytes, &option)
	if err != nil {
		fmt.Printf("Error unmarshalling essentialInfo.json: %v\n", err)
		os.Exit(1)
	}
	return essential, option
}
