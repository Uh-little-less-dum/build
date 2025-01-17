package build_test

import (
	"path/filepath"

	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
	ulld_plugin "github.com/Uh-little-less-dum/build/pkg/plugin"
	target_paths "github.com/Uh-little-less-dum/build/pkg/targetPaths"
	ulld_test "github.com/Uh-little-less-dum/go-utils/pkg/testing"
)

func GetLocalBuildManager() *build_config.BuildManager {
	b := build_config.GetBuildManager()
	b.Init([]string{})
	b.SetAppConfigPath(ulld_test.GetLocalAppConfigPath())
	return b
}

func ApplyLocalPlugins(b *build_config.BuildManager) {
	pathMap := ulld_test.GetLocalPluginPathMap()
	for _, p := range b.Plugins {
		p.SetInstallLocation(filepath.Dir(pathMap[p.Name]))
	}
}

// TEST: Enable these methods once the slotMap struct is in order.
func generatePageConflict(b *build_config.BuildManager) {
}

func generateSlotConflict(b *build_config.BuildManager) {
}

func GetMockPlugins(withPageConflict, withSlotConflict bool) []*ulld_plugin.Plugin {
	b := GetLocalBuildManager()
	targetDir := ulld_test.TestOutputRoot()
	res := b.AppConfig().GatherPlugins(target_paths.NewTargetPaths(targetDir))
	if withPageConflict {
		generatePageConflict(b)
	}
	if withSlotConflict {
		generateSlotConflict(b)
	}
	return res
}
