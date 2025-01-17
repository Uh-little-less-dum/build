package stage_pre_conflict_resolve_build

import (
	"fmt"

	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
	file_handlers_package_json "github.com/Uh-little-less-dum/go-utils/pkg/buildFiles/file_handlers/packageJsonHandler"
	"github.com/charmbracelet/log"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func SetInitialPackageJsonData(cfg *build_config.BuildManager) {
	b := cfg.RootPackageJson.Bytes()
	j := gjson.ParseBytes(b)
	var err error
	for d := range file_handlers_package_json.DependencyTypes() {
		data := j.Get(fmt.Sprintf("ulld-%s", d))
		if data.Exists() {
			b, err = sjson.SetBytes(b, d, data)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	cfg.RootPackageJson.SetBytes(b)
}
