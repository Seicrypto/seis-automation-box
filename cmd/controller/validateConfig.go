package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"seis-automation-box/internal/envStaffs"
)

func isValidUserName(name string) bool {
	re := regexp.MustCompile(`^[a-zA-Z\s\-]+$`)
	return re.MatchString(name)
}

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
func essentialSettingCheck(settingFilePath string) error {

	data := envStaffs.LoadEssentialInfoConfig(settingFilePath)
	// Verify userName.
	if !isValidUserName(data["userName"]) {
		return fmt.Errorf("invalid userName format")
	}

	// Verify userEmail.
	if !isValidEmail(data["userEmail"]) {
		return fmt.Errorf("invalid userEmail format")
	}

	return nil
}

func splitPascalCase(input string) []string {
	re := regexp.MustCompile(`[A-Z][^A-Z]*`)
	parts := re.FindAllString(input, -1)
	return parts
}

func isValidOptionConfig(data []byte) (bool, error) {
	var config map[string]map[string]bool
	if err := json.Unmarshal(data, &config); err != nil {
		return false, err
	}
	return true, nil
}

func optionSettingCheck(settingFilePath string) error {

	data, err := os.ReadFile(settingFilePath)
	if err != nil {
		fmt.Println("Error reading optionInst.json:", err)
		os.Exit(1)
	}
	// Verify option setting structure.
	valid, err := isValidOptionConfig(data)
	if err != nil {
		return err
	}

	if !valid {
		return fmt.Errorf("Invalid option config")
	}
	// Verify optional feature funcs equal option setting map.
	// Haven't build.
	return nil
}

// Get the executing path.
func execPath() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Failed to get executable path: %v", err)
	}
	exeDir := filepath.Dir(exePath)
	return exeDir
}

// Plus the configs path.
func essentialInfoConfigPath() string {
	configPath := filepath.Join(execPath(), "configs", "essentialInfo.json")
	return configPath
}

func optionConfigPath() string {
	configPath := filepath.Join(execPath(), "configs", "option.json")
	return configPath
}

func ValidateConfig() {
	if err := essentialSettingCheck(essentialInfoConfigPath()); err != nil {
		fmt.Fprintf(os.Stderr, "EssentialInfo configuration file validation fails: %v\n", err)
		os.Exit(1)
	}
	if err := optionSettingCheck(optionConfigPath()); err != nil {
		fmt.Fprintf(os.Stderr, "Option configuration file validation fails: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Configs format verified.")
}
