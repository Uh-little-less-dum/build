package ulld_plugin

import (
	"cmp"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	conflicts_page "github.com/Uh-little-less-dum/build/pkg/conflicts/page"
	conflicts_slot "github.com/Uh-little-less-dum/build/pkg/conflicts/slot"
	slot_map "github.com/Uh-little-less-dum/build/pkg/slotMap"
	target_paths "github.com/Uh-little-less-dum/build/pkg/targetPaths"
	"github.com/charmbracelet/log"
	"github.com/tidwall/gjson"
)

type Plugin struct {
	Name            string
	Version         string
	Slot            slot_map.SlotKey
	pluginConfig    PluginConfig
	installLocation string
	paths           *target_paths.TargetPaths
	// Map input urls to urls created by the user durng the conflict resolution stage.
	ModifiedUrlMap map[string]string
}

// Returns the install string for a specific package.
// Example: redux or react-redux@2.15.21
func (plugin *Plugin) InstallString() string {
	buildLocal := os.Getenv("ULLD_LOCAL_BUILD")
	if (buildLocal == "true") && (strings.HasPrefix(plugin.Name, "@ulld/")) {
		// currentPackages := goutils
		// TODO: Implement this to work on the build locally.
		log.Fatal("Need to handle the implementation of local file paths here.")
		// return fmt.Sprintf("%s@file://")
	}
	if plugin.Version == "latest" {
		return plugin.Name
	}
	return fmt.Sprintf("%s@%s", plugin.Name, plugin.Version)
}

// WARN: This is likely to be unreliable for some advanced use case. This will either need to dynamically check for pnpm global stores or disable them all together via the generated npmrc.
func (p Plugin) InstallLocation() string {
	return filepath.Join(p.paths.TargetDir(), cmp.Or(p.installLocation, filepath.Join("node_modules", p.Name)))
}

func (p *Plugin) SetInstallLocation(path string) {
	p.installLocation = path
}

func (p *Plugin) Config() gjson.Result {
	return p.pluginConfig.Config(p.InstallLocation(), p.Name)
}

func (p *Plugin) Events() PluginEvents {
	return *NewPluginEvents(&p.pluginConfig)
}

func (p *Plugin) Components() []*PluginComponent {
	var res []*PluginComponent
	slot := p.Slot
	components := p.Config().Get("components").Array()
	for _, c := range components {
		res = append(res, NewPluginComponent(c, slot))
	}
	return res
}

func (p *Plugin) Pages() []*PluginPage {
	var res []*PluginPage
	pages := p.Config().Get("pages").Array()
	for _, data := range pages {
		res = append(res, NewPluginPage(data))
	}
	return res
}

func (p *Plugin) Parsers() []*PluginParser {
	var res []*PluginParser
	p.Config().Get("parsers").ForEach(func(key, value gjson.Result) bool {
		res = append(res, NewPluginParser(key, value))
		return true
	})
	return res
}

func (p *Plugin) SettingsPage() *PluginSettings {
	return NewPluginSettings(&p.pluginConfig.data)
}

func (p *Plugin) AllEmbeddables() []*PluginEmbeddable {
	var res []*PluginEmbeddable
	components := p.Components()
	for _, item := range components {
		res = append(res, item.Embeddables()...)
	}
	return res
}

// TEST: These conflict gatherng methods need to be tested much more thoroughly when on wifi and power, and after the build is in working order.
func (b *Plugin) HasPageConflict(p *Plugin) []*conflicts_page.Conflict {
	var res []*conflicts_page.Conflict
	for _, page1 := range b.Pages() {
		targetUrl := page1.data.Get("targetUrl").Str
		for _, page2 := range p.Pages() {
			if targetUrl == page2.data.Get("targetUrl").Str {
				res = append(res, conflicts_page.NewPageConflict(targetUrl))
			}
		}
	}
	return res
}

func (b *Plugin) HasSlotConflict(p *Plugin) []*conflicts_slot.Conflict {
	var res []*conflicts_slot.Conflict
	for _, c1 := range b.Components() {
		slot1 := c1.data.Get("slot").Str
		if slot1 != "" {
			for _, c2 := range p.Components() {
				if slot1 == c2.data.Get("slot").Str {
					res = append(res, conflicts_slot.NewSlotConflict(slot1))
				}
			}
		}
	}
	return res
}
