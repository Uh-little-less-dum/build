package stage_pre_conflict_resolve_build

import (
	"os"

	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
	"github.com/charmbracelet/log"
)

// TODO: This doesn't seem to ever resolve. This is likely an issue with the buildStream model itself, and not necessarily this method.
func InstallDependencies(cfg *build_config.BuildManager) {
	if os.Getenv("ULLD_LOCAL_DEV") == "true" {
		return
	}
	packageManager := cfg.PackageManager()
	c := packageManager.Add(cfg.Installables())
	c.Dir = cfg.TargetDir()
	err := c.Run()
	if err != nil {
		log.Fatal(err)
	}
}
