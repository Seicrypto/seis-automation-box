package optionalfeatures

import (
	"encoding/json"
	"fmt"
	"os"
)

type OptionConfig map[string]map[string]bool

func LoadOptionConfig() OptionConfig {

	var option OptionConfig

	optionBytes, err := os.ReadFile("./configs/option.json")
	if err != nil {
		fmt.Println("Error reading optionInst.json:", err)
		os.Exit(1)
	}
	err = json.Unmarshal(optionBytes, &option)
	if err != nil {
		fmt.Printf("Error unmarshalling essentialInfo.json: %v\n", err)
		os.Exit(1)
	}

	return option
}
