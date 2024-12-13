package utils_test

import (
	"testing"

	"github.com/Uh-little-less-dum/build/pkg/utils"
)

func Test_IsInstalled(t *testing.T) {
	var vals = []struct {
		name     string
		inputVal string
		expected bool
	}{
		{"IsInstalled returns true if package is available", "tsx", true},
		{"IsInstalled returns false if package is unavailable", "stripyDipy", false},
	}
	for _, tt := range vals {
		t.Run(tt.name, func(t *testing.T) {
			b := utils.IsInstalled(tt.inputVal)
			if b != tt.expected {
				t.Errorf("Expected '%v', received '%v'", tt.expected, b)
			}
		})
	}
}
