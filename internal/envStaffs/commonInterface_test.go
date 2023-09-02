package envStaffs_test

import "os/exec"

type MockCmdRunner struct {
	errToReturn error
}

// Just do mock to test the logic, not succeseed or not in real envirement.
func (m *MockCmdRunner) Run(cmd *exec.Cmd) error {
	return m.errToReturn
}
