package post_conflict_resolve_stages

import (
	"sync"

	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
)

// RESUME: Need to handle all of these methods, and then return to /Users/bigsexy/Desktop/Go/projects/ulld/build/pkg/buildScript/stages/post_conflict_resolve_build/subStageTree.go and keep moving down the list.
// Writes the plugin output for the followng types:
//   - [x] GeneratedMarkdown
//   - [x] Additional Page
//   - [x] Settings Page
//   - [x] Parser methods
//   - [ ] Tailwind wrapper methods
//   - [x] Scss outputs (handled in the WriteStyleSheets subStage)
func WritePluginOutput(cfg *build_config.BuildManager) {
	var wg sync.WaitGroup
	for _, p := range cfg.Plugins {
		for _, page := range p.Pages() {
			wg.Add(1)
			go func() {
				defer wg.Done()
				page.WriteOutput(cfg.Paths)
			}()
		}
		for _, c := range p.Components() {
			c.WriteDocsOutput(&wg)
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			p.SettingsPage().WriteOutput(cfg.Paths)
		}()
	}
	wg.Wait()
}
