package stage_gather_build_config

import (
	"path/filepath"

	fs_utils "github.com/Uh-little-less-dum/build/pkg/fs"
	app_config "github.com/Uh-little-less-dum/go-utils/pkg/buildFiles/appConfig"
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

func GatherAppConfig() *app_config.AppConfig {
	v := viper.GetViper().GetString("targetDir")
	p := filepath.Join()
	if v == "" {
		p := filepath.Join(v, "appConfig.ulld.json")
		if !fs_utils.Exists(p) {
			log.Fatal("Cannot gather the target directory successfully.")
		}
	}
	return app_config.NewAppConfig(p)
}
