package stage_transpile_app

import (
	package_managers "github.com/Uh-little-less-dum/build/pkg/packageManager"
	"github.com/charmbracelet/log"
)

func Run(packageManager package_managers.PackageManager, targetDir string) {
	c := packageManager.RunScript("next", "build")
	c.Dir = targetDir
	err := c.Run()
	if err != nil {
		log.Fatal(err)
	}
}
