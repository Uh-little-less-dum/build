package stage_install_dependencies

import (
	"log"
	"os"
	"os/exec"

	"github.com/spf13/viper"
)

func InstallDependencies() {
	rootDir := viper.GetViper().GetString("targetDir")
	packageManager := viper.GetViper().GetString("packageManager")
	cmd := exec.Command(packageManager, "install")
	cmd.Stdout = os.Stdout
	if rootDir != "" {
		cmd.Dir = rootDir
	}
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
