package ulld_plugin_test

import (
	"math/rand"
	"path/filepath"

	"github.com/Uh-little-less-dum/build/mocks"
	ulld_plugin "github.com/Uh-little-less-dum/build/pkg/plugin"
	slot_map "github.com/Uh-little-less-dum/build/pkg/slotMap"
	ulld_test "github.com/Uh-little-less-dum/go-utils/pkg/testing"
)

func MockLocalPlugin(idx int) *ulld_plugin.Plugin {
	paths := ulld_test.GetLocalPluginConfigPaths()
	if idx == -1 {
		idx = rand.Intn(len(paths))
	}
	item := paths[idx]
	packagesDir := filepath.Join(ulld_test.DevRoot(), "packages")
	installLoc := filepath.Dir(item)
	// subDir, err := filepath.Rel(packagesDir, item)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// subDir = strings.ReplaceAll(subDir, "/pluginConfig.ulld.json", "")
	plugin := ulld_plugin.NewPluginWithInstallLocation(installLoc, "1.0.0", slot_map.Bibliography, mocks.TargetPaths(), packagesDir)
	return plugin
}
