package post_conflict_resolve_stages

import (
	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
	parsers "github.com/Uh-little-less-dum/build/pkg/parsers"
)

func WriteAdditionalParserLists(cfg *build_config.BuildManager) {
	pl := parsers.NewParserLists(cfg)
	pl.WriteOutput(cfg)
}
