package stage_pre_conflict_resolve_build

import (
	"log"
	"path/filepath"

	file_handlers_package_json "github.com/Uh-little-less-dum/go-utils/pkg/buildFiles/file_handlers/packageJsonHandler"
)

func GatherRootPackageJson(targetDir string) file_handlers_package_json.PackageJsonHandler {
	handler := file_handlers_package_json.NewPackageJsonHandler(filepath.Join(targetDir, "package.json"))
	if !handler.File.Exists() {
		log.Fatal("Could not locate the root package.json file. Cannot continue with the build.")
	}
	return handler
}
