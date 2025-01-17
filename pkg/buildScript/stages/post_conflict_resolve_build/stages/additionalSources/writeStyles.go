package additional_sources_stages

import build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"

func WriteAdditionalSourcesBasedStyles(cfg *build_config.BuildManager) {
	cfg.AdditionalSources.WriteScssFiles(cfg.Paths)
}
