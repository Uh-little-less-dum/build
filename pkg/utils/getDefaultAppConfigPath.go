package utils

import (
	"os"
	"path/filepath"

	env_vars "github.com/Uh-little-less-dum/build/pkg/envVars"
	fs_utils "github.com/Uh-little-less-dum/build/pkg/fs"
	"github.com/charmbracelet/log"
)

func GetDefaultAppConfigPath() string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	cwdRelPath := filepath.Join(cwd, "appConfig.ulld.json")
	if fs_utils.Exists(cwdRelPath) {
		return cwdRelPath
	}
	additionalSourcesPath := os.Getenv(string(env_vars.AdditionalSources))
	if additionalSourcesPath == "" {
		return ""
	}
	additionalSourcesPath = filepath.Join(additionalSourcesPath, "appConfig.ulld.json")

	if fs_utils.Exists(additionalSourcesPath) {
		return additionalSourcesPath
	} else {
		return ""
	}
}
