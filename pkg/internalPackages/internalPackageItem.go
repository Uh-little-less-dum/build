package internal_package_items

import (
	"fmt"

	file_handlers_package_json "github.com/Uh-little-less-dum/go-utils/pkg/buildFiles/file_handlers/packageJsonHandler"
)

type InternalPackageItem struct {
	Name    string
	Version string
}

// Replaces version of this internal package in the provided PackageJsonFile with the latest version.
func (i InternalPackageItem) ApplyVersion(item *file_handlers_package_json.PackageJsonHandler) {
	data := item.Json()
	for d := range file_handlers_package_json.DependencyTypes() {
		valPath := fmt.Sprintf("%s.%s", d, i.Name)
		r := data.Get(valPath).Str
		if r != "" {
			item.Set(valPath, i.Version)
		}
	}
}
