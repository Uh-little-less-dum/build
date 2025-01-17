package post_conflict_resolve_stages

import build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"

func WriteStyleSheets(cfg *build_config.BuildManager) {
	cfg.AdditionalSources.ClearGeneratedStyles(cfg.Paths)
	cfg.Styles.WriteScssFiles(cfg.Paths, cfg.GetCopyStyleSheets())
}
