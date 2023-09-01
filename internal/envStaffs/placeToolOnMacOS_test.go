package envStaffs_test

import (
	"os/exec"
	"seis-automation-box/internal/envStaffs"
	"testing"
)

type MockCmdRunner struct {
	errToReturn error
}

func (m *MockCmdRunner) Run(cmd *exec.Cmd) error {
	return m.errToReturn
}

func TestInstallHomebrew(t *testing.T) {
	mockRunner := &MockCmdRunner{
		errToReturn: nil,
	}

	err := envStaffs.InstallHomebrew(mockRunner)
	if err != nil {
		t.Fatalf("expected no error, but got: %v", err)
	}
}
