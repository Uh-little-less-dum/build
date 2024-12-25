package stage_pre_conflict_resolve_build_test

import (
	"testing"

	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
	stage_pre_conflict_resolve_build "github.com/Uh-little-less-dum/build/pkg/buildScript/stages/pre_conflict_resolve_build"
	ulld_plugin "github.com/Uh-little-less-dum/build/pkg/plugin"
)

func iWantToSeeYourManager() *build_config.BuildManager {
	b := build_config.GetInitialBuildManager()

	b.Init([]string{})
	b.AppConfigPath = "/Users/bigsexy/dev-utils/ulld/appConfig.ulld.json"
	b.Plugins = []*ulld_plugin.Plugin{}
	return b
}

func Test_GatherPlugins(t *testing.T) {
	t.Run("Gathered plugins from appConfig", func(t *testing.T) {
		b := iWantToSeeYourManager()
		stage_pre_conflict_resolve_build.GatherPlugins(b)
		if len(b.Plugins) == 0 {
			t.Log("Failed to find any plugins.")
			t.Fail()
		}
	})
}
