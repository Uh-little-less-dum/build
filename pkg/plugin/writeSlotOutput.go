package ulld_plugin

import (
	"fmt"
	"sync"

	target_paths "github.com/Uh-little-less-dum/build/pkg/targetPaths"
	"github.com/tidwall/gjson"
)

func (p *Plugin) FormattedExport(subPath string) string {
	return fmt.Sprintf("%s/%s", p.Name, subPath)
}

func (p *Plugin) WriteSlotOutput(slotMapData gjson.Result, paths target_paths.TargetPaths, wg *sync.WaitGroup) {
	if !p.HasSlot() {
		return
	}
	components := p.Components()
	wg.Add(len(components))
	for _, component := range components {
		go func() {
			defer wg.Done()
			component.WriteSlotOutput(slotMapData, p.Config(), component.Data(), p.Name, paths)
		}()
	}
	pages := p.Pages()
	for _, page := range pages {
		if page.IsSlotPage(slotMapData) {
			wg.Add(1)
			go func() {
				defer wg.Done()
				page.WriteSlotOutput(slotMapData, p.Name, paths)
			}()
		}
	}
}
