package envStaffs

import "os/exec"

type CmdRunner interface {
	Run(cmd *exec.Cmd) error
}

type RealCmdRunner struct{}

func (r *RealCmdRunner) Run(cmd *exec.Cmd) error {
	return cmd.Run()
}
