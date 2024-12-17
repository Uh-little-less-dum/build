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
		{"IsInstalled returns true if package is available", "tsx", false},
		{"IsInstalled returns false if package is unavailable", "stripyDipy", false},
	}
	for _, tt := range vals {
		t.Run(tt.name, func(t *testing.T) {
			_, ok := utils.IsInstalled(tt.inputVal)
			if ok != tt.expected {
				t.Logf("Expected %v, found %v for package %s", ok, tt.expected, tt.inputVal)
				t.Fail()
			}
		})
	}
}
