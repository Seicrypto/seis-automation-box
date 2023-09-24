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
	cmd.Stdin = os.Stdin   // Let command stdin as go project stdin
	cmd.Stdout = os.Stdout // Let command stdout as go project stdout
	cmd.Stderr = os.Stderr // Let command stderr as go project stderr
	return cmd.Run()
}

// Load essentailInfo.json.
type UserConfig map[string]string

func LoadEssentialInfoConfig(essentailInfoPath string) UserConfig {

	var essential UserConfig

	essentialBytes, err := os.ReadFile(essentailInfoPath)
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
