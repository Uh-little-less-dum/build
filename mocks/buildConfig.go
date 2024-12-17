package mocks

import (
	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
	"github.com/spf13/afero"
)

func GetMockBuildConfig() *build_config.BuildManager {
	b := build_config.GetBuildManager()
	d := afero.NewMemMapFs()
	b.Fs = &d
	return b
}
