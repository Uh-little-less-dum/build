package stage_pre_conflict_resolve_build

import (
	"log"
	"path/filepath"

	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
	file_handlers_package_json "github.com/Uh-little-less-dum/go-utils/pkg/buildFiles/file_handlers/packageJsonHandler"
)

func GatherRootPackageJson(cfg *build_config.BuildManager) {
	handler := file_handlers_package_json.NewPackageJsonHandler(filepath.Join(cfg.TargetDir(), "package.json"))
	if !handler.File.Exists() {
		log.Fatal("Could not locate the root package.json file. Cannot continue with the build.")
	}
	cfg.SetRootPackageJson(handler)
}
