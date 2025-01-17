package post_conflict_resolve_stages

import (
	"sync"

	event_handler_types "github.com/Uh-little-less-dum/build/pkg/buildConstants/eventTypes"
	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
)

// Writes onBackup handlers as a single function to the target file.
func WriteEventHandlers(cfg *build_config.BuildManager) {
	var wg sync.WaitGroup
	eventTypes := event_handler_types.PluginSyncTypes()
	for _, item := range eventTypes {
		wg.Add(1)
		go func() {
			d := cfg.EventHandlerListOfType(item)
			d.WriteOutput(cfg.Paths)
			defer wg.Done()
		}()
	}
	wg.Wait()
}
