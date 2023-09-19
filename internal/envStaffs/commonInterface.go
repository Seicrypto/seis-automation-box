package envStaffs

import "os/exec"

type CmdRunner interface {
	Run(cmd *exec.Cmd) error
}

type RealCmdRunner struct{}

func (r *RealCmdRunner) Run(cmd *exec.Cmd) error {
	return cmd.Run()
}

// Config should be initiate from cmd/controller/loadConfig.go
type UserConfig map[string]string

type OptionConfig map[string]map[string]bool
