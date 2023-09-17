package envStaffs

import (
	"fmt"
	"os/exec"
)

// Set up gitconfig:
// global username
func GitConfigUserName(runner CmdRunner, userName string) error {
	cmdGitUserName := `user.name`
	cmdGitUserNameContent := fmt.Sprintf(`"%s"`, userName)
	gitConfigUserName := exec.Command("git", "config", "--global", cmdGitUserName, cmdGitUserNameContent)
	return runner.Run(gitConfigUserName)
}

func setGitconfig() {
	// Set up gitconfig.username

	// Set up gitconfig.usermail

	// Set up CR LF

	// Set up diff tool
}

// Set up .ZSHrc
func setZSHrc() {
	// Set up plugin

	// Set the theme

}

func SettingToolOnMacOS() {
	setGitconfig()
	setZSHrc()
}
