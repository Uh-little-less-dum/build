package post_conflict_resolve_stages

import (
	"sync"

	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
)

func WriteSlotOutput(cfg *build_config.BuildManager) {
	var wg sync.WaitGroup
	slotMapData := cfg.SlotMapData()
	for _, plugin := range cfg.Plugins {
		wg.Add(1)
		go func() {
			defer wg.Done()
			plugin.WriteSlotOutput(slotMapData, *cfg.Paths, &wg)
		}()
	}
	wg.Wait()
}
