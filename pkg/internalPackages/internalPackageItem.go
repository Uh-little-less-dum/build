package internal_package_items

import (
	file_handlers_package_json "github.com/Uh-little-less-dum/go-utils/pkg/buildFiles/file_handlers/packageJsonHandler"
	"github.com/charmbracelet/log"
)

type InternalPackageItem struct {
	Name    string
	Version string
}

// Replaces version of this internal package in the provided PackageJsonFile with the latest version.
func (i InternalPackageItem) ApplyVersion(item *file_handlers_package_json.PackageJsonHandler) {
	data, err := item.GetData()
	if err != nil {
		log.Fatal(err)
	}
	log.Info(data)
	// BUG: Fix this now that everything has been cleaned up.
	// Regular dependencies
	// _, ok := data.Dependencies[i.Name]
	// if ok {
	// 	data.Data.Dependencies[i.Name] = i.Version
	// }
	// // Peer dependencies
	// _, ok = data.Data.PeerDependencies[i.Name]
	// if ok {
	// 	data.Data.PeerDependencies[i.Name] = i.Version
	// }
	// // Dev dependencies
	// _, ok = data.Data.DevDependencies[i.Name]
	// if ok {
	// 	data.Data.DevDependencies[i.Name] = i.Version
	// }
}
