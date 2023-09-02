package envStaffs_test

import (
	"seis-automation-box/internal/envStaffs"
	"testing"
)

func TestInstallHomebrew(t *testing.T) {
	mockRunner := &MockCmdRunner{
		errToReturn: nil,
	}

	err := envStaffs.InstallHomebrew(mockRunner)
	if err != nil {
		t.Fatalf("expected no error, but got: %v", err)
	}
}

func TestBrewInstGit(t *testing.T) {
	mockRunner := &MockCmdRunner{
		errToReturn: nil,
	}

	err := envStaffs.BrewInstGit(mockRunner)
	if err != nil {
		t.Fatalf("expected no error, but got: %v", err)
	}
}

func TestCmdGitCredientail(t *testing.T) {
	mockRunner := &MockCmdRunner{
		errToReturn: nil,
	}

	err := envStaffs.CmdGitCredential(mockRunner)
	if err != nil {
		t.Fatalf("expected no error, but got: %v", err)
	}
}

func TestBrewInstITerm2(t *testing.T) {
	mockRunner := &MockCmdRunner{
		errToReturn: nil,
	}

	err := envStaffs.BrewInstITerm2(mockRunner)
	if err != nil {
		t.Fatalf("expected no error, but got: %v", err)
	}
}
func TestBrewInstZSH(t *testing.T) {
	mockRunner := &MockCmdRunner{
		errToReturn: nil,
	}

	err := envStaffs.BrewInstZSH(mockRunner)
	if err != nil {
		t.Fatalf("expected no error, but got: %v", err)
	}
}
func TestInstVSCode(t *testing.T) {
	mockRunner := &MockCmdRunner{
		errToReturn: nil,
	}

	err := envStaffs.InstVSCode(mockRunner)
	if err != nil {
		t.Fatalf("expected no error, but got: %v", err)
	}
}
