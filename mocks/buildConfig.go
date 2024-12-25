package mocks

import (
	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
	target_paths "github.com/Uh-little-less-dum/build/pkg/targetPaths"
	ulld_test "github.com/Uh-little-less-dum/go-utils/pkg/testing"
	"github.com/spf13/afero"
)

func MockBuildConfig() *build_config.BuildManager {
	b := build_config.GetBuildManager()
	b.SetTargetDir(ulld_test.TestOutputRoot())
	d := afero.NewMemMapFs()
	b.Fs = &d
	return b
}

func TargetPaths() *target_paths.TargetPaths {
	return target_paths.NewTargetPaths(ulld_test.TestOutputRoot())
}
