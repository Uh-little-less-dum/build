package post_conflict_resolve_stages

import (
	"sync"

	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
	parser_list "github.com/Uh-little-less-dum/go-utils/pkg/buildFiles/parserList"
)

// Writes the plugin output for the followng types:
//   - [x] GeneratedMarkdown
//   - [x] Additional Page
//   - [x] Settings Page
//   - [x] Parser methods
//   - [ ] Tailwind wrapper methods
//   - [ ] Scss outputs
func WritePluginOutput(cfg *build_config.BuildManager) {
	var wg sync.WaitGroup
	for _, p := range cfg.Plugins {
		for _, page := range p.Pages() {
			wg.Add(1)
			go func() {
				defer wg.Done()
				page.WriteOutput()
			}()
		}
		for _, c := range p.Components() {
			wg.Add(1)
			go func() {
				defer wg.Done()
				dp, dpFull := c.DocPaths()
				if dp != "" {
					c.WriteDocsOutput(dp, false)
				}
				if dpFull != "" {
					c.WriteDocsOutput(dpFull, true)
				}
			}()
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			p.SettingsPage().WriteOutput()
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			pl := parser_list.NewParserLists(cfg)
			pl.WriteOutput(cfg)
		}()
	}
	wg.Wait()
}
