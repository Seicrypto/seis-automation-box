package envStaffs

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
)

type CmdRunner interface {
	Run(cmd *exec.Cmd) error
}

type RealCmdRunner struct{}

func (r *RealCmdRunner) Run(cmd *exec.Cmd) error {
	return cmd.Run()
}

// Load essentailInfo.json.
type UserConfig map[string]string

func LoadEssentialInfoConfig() UserConfig {

	var essential UserConfig

	essentialBytes, err := os.ReadFile("./configs/essentialInfo.json")
	if err != nil {
		fmt.Println("Error reading essentialInfo.json:", err)
		os.Exit(1)
	}
	err = json.Unmarshal(essentialBytes, &essential)
	if err != nil {
		fmt.Printf("Error unmarshalling essentialInfo.json: %v\n", err)
		os.Exit(1)
	}

	return essential
}
