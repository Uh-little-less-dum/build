package stage_pre_conflict_resolve_build

import build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"

func PackageManagerModifyPackageJson(cfg *build_config.BuildManager) {
	if cfg.PackageManager().ModifiesPackageJson() {
		cfg.RootPackageJson.SetBytes(cfg.PackageManager().ModifyPackageJson(cfg.RootPackageJson.Bytes()))
	}
}
