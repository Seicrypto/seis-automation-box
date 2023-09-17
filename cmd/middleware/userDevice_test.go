package middleware_test

import (
	"seis-automation-box/cmd/middleware"
	"testing"
)

func TestCurrentDevice(t *testing.T) {
	funExpect := []string{
		"darwin",
		"linux",
		"windows",
	}
	valid := false
	funcResult := middleware.CurrentDevice()
	for _, content := range funExpect {
		if funcResult == content {
			valid = true
			break
		}
		if !valid {
			t.Errorf("Unexpect iso system: %s", funcResult)
		}
	}
}
