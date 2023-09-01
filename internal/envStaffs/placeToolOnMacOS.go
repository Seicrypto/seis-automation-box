package envStaffs

import (
	"errors"
	"fmt"
	"log"
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

// Check and install "Homebrew" on envirement.
func installHomebrew(runner CmdRunner) error {
	cmdStr := `curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh`
	cmdInstallHomebrew := exec.Command("/bin/bash", "-c", cmdStr)
	return runner.Run(cmdInstallHomebrew)
}

// Brew install git and git CLI. (Let brew package check install or not.)
func brewInstGit(runner CmdRunner) error {
	cmdBrewGit := exec.Command("brew", "install", "git")
	return runner.Run(cmdBrewGit)
}

// Install git credential-osxkeychain. (For storing crediential.)

func cmdGitCredientail(runner CmdRunner) error {
	cmdContent := "crediential.helpler"
	cmdGitCrediential := exec.Command("git", "config", "--global", cmdContent, "osxkeychan")
	return runner.Run(cmdGitCrediential)
}

// Install iTerm2
func brewInstITerm2(runner CmdRunner) error {
	cmdBrewITerm2 := exec.Command("brew", "install", "--cask", "iterm2")
	return runner.Run(cmdBrewITerm2)
}

// Install ZSH.
func brewInstZSH(runner CmdRunner) error {
	cmdBrewZSH := exec.Command("brew", "install", "zsh")
	return runner.Run(cmdBrewZSH)
}

// Install Fonts funcs:
func getAvailableDirName(base string) (string, error) {
	dirName := base
	i := 1
	for {
		if _, err := os.Stat(dirName); os.IsNotExist(err) {
			// Directory doesn't exist, we can use this name
			return dirName, nil
		}
		// Directory exists, try next name
		dirName = fmt.Sprintf("%s-%d", base, i)
		i++
		// Safety check to avoid infinite loops, adjust as necessary
		if i > 1000 {
			return "", errors.New("too many directories with similar names")
		}
	}
}

func cloneFonts(runner CmdRunner, targetDir string) error {
	repoURL := `https://github.com/powerline/fonts.git`
	cmdCloneFonts := exec.Command("git", "clone", repoURL, targetDir, "--depth=1")
	return runner.Run(cmdCloneFonts)
}

func installFonts(runner CmdRunner, dir string) error {
	cmd := exec.Command("./install.sh")
	cmd.Dir = dir
	return runner.Run(cmd)
}

func rmDir(runner CmdRunner, dir string) error {
	cmd := exec.Command("rm", "-rf", dir)
	return runner.Run(cmd)
}

func placeFonts() {
	runner := &RealCmdRunner{}

	dir, err := getAvailableDirName("fonts")
	if err != nil {
		log.Fatal(err)
	}

	if err := cloneFonts(runner, dir); err != nil {
		log.Fatal(err)
	}
	if err := installFonts(runner, dir); err != nil {
		log.Fatal(err)
	}
	if err := rmDir(runner, dir); err != nil {
		log.Fatal(err)
	}
}

// Install VSCode and code CLI (Issue: how to make sure the VSCode GUI already install.)
func instVSCode(runner CmdRunner) error {
	instTarget := `visual-studio-code`
	cmdInstVSCode := exec.Command("brew", "install", "--cask", instTarget)
	return runner.Run(cmdInstVSCode)
}

// Install VSCode extensions.
func installVSCodeExtension() {
	// Install extension:
	// 1. Ayu theme
	// 2. Markdown lint / preview
	// 3. Markdown mermaid preview
	// 4. Prettier -Code formatter
	// 5. Material Icon
	// 6. Git Graph

}

// Main
func PlaceToolOnMac() {
	runner := &RealCmdRunner{}
	if err := installHomebrew(runner); err != nil {
		log.Fatal(err)
	}
	if err := brewInstGit(runner); err != nil {
		log.Fatal(err)
	}
	if err := brewInstITerm2(runner); err != nil {
		log.Fatal(err)
	}
	if err := brewInstZSH(runner); err != nil {
		log.Fatal(err)
	}
	placeFonts()
	if err := instVSCode(runner); err != nil {
		log.Fatal(err)
	}
}
