package post_conflict_resolve_stages

import (
	"sync"

	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
)

func WriteNoteTypeData(cfg *build_config.BuildManager) {
	nt := cfg.AppConfig().NoteTypes()
	var wg sync.WaitGroup
	for _, n := range nt {
		wg.Add(1)
		go func() {
			defer wg.Done()
			n.WriteOutput(cfg)
		}()
	}
	wg.Wait()
}
