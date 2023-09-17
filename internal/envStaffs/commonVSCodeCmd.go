package envStaffs

import "os/exec"

// This part is going to establish in new repo with bash and powershell.
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

// Set up VSCode json setting.
func setVSCodeJson() {
	// Set ZSH as default terminal.
	// Set up fonts and size.
}

// Install VSCode Extension:
// Ayu
func InstVSCodeAyu(runner CmdRunner) error {
	cmdOp := `--install-extension`
	exAyu := `teabyii.ayu`
	instVSCodeExAyu := exec.Command("code", cmdOp, exAyu)
	return runner.Run(instVSCodeExAyu)
}
