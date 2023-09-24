package terminal

import (
	"os"
	"os/exec"
)

// This func is for testing cmd in terminal.
func StandardTerminalTest() {
	cmd := exec.Command("bash", "-c", "read -p 'Enter your name: ' name; echo Hello, $name!")
	cmd.Stdin = os.Stdin   // Let command stdin as go project stdin
	cmd.Stdout = os.Stdout // Let command stdout as go project stdout
	cmd.Stderr = os.Stderr // Let command stderr as go project stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
