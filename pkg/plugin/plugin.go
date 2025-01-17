package ulld_plugin

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	conflicts_page "github.com/Uh-little-less-dum/build/pkg/conflicts/page"
	conflicts_slot "github.com/Uh-little-less-dum/build/pkg/conflicts/slot"
	navigation_link "github.com/Uh-little-less-dum/build/pkg/plugin/navigationLink"
	plugin_setting_page_data "github.com/Uh-little-less-dum/build/pkg/plugin/settingPageData"
	slot_map "github.com/Uh-little-less-dum/build/pkg/slotMap"
	target_paths "github.com/Uh-little-less-dum/build/pkg/targetPaths"
	file_handlers_package_json "github.com/Uh-little-less-dum/go-utils/pkg/buildFiles/file_handlers/packageJsonHandler"
	copy_file "github.com/Uh-little-less-dum/go-utils/pkg/fs/copyFile"
	"github.com/charmbracelet/log"
	"github.com/tidwall/gjson"
)

type Plugin struct {
	Name         string
	Version      string
	Slot         slot_map.SlotKey
	pluginConfig PluginConfig
	// installLocation points to the directory containing the plugins package.json directly, not the node_modules folder. If set to an empty string, the <target dir>/node_modules/<pluginName> path will be used.
	installLocation string
	paths           *target_paths.TargetPaths
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
	if p.installLocation != "" {
		return p.installLocation
	}
	return filepath.Join(p.paths.TargetDir(), filepath.Join("node_modules", p.Name))
}

func (p *Plugin) SetInstallLocation(path string) {
	p.installLocation = path
}

func (p *Plugin) Config() gjson.Result {
	return p.pluginConfig.Config(p.InstallLocation(), p.Name)
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
	i := 1
	p.Config().Get("parsers").ForEach(func(key, value gjson.Result) bool {
		res = append(res, NewPluginParser(key, value, p.Name, i))
		return true
	})
	return res
}

func (p Plugin) NavigationLinks(c chan navigation_link.NavigationLink) {
	data := p.Config().Get("navigationLinks")
	if data.Exists() {
		for _, itemData := range data.Array() {
			c <- navigation_link.NewNavigationLink(itemData)
		}
	}
}

func (p *Plugin) SettingsPage() *PluginSettings {
	return NewPluginSettings(&p.pluginConfig.data)
}

// TODO: Make sure to gather this data properly on the typescript side of things. That page isn't yet built.
func (p *Plugin) SettingsPageUrl() string {
	return fmt.Sprintf("/settings/%s", p.Name)
}

func (p *Plugin) AllEmbeddables() []PluginEmbeddableTemplateStruct {
	var res []PluginEmbeddableTemplateStruct
	components := p.Components()
	for i, item := range components {
		res = append(res, item.Embeddables(p.Name, item.ComponentName(), item.ExportPath(), i)...)
	}
	return res
}

func (p *Plugin) ShouldTranspile() bool {
	return p.pluginConfig.data.Get("transpile").Bool()
}

func (b *Plugin) PackageJson() file_handlers_package_json.PackageJsonHandler {
	return file_handlers_package_json.NewPackageJsonHandler(filepath.Join(b.InstallLocation(), "package.json"))
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

func (b *Plugin) ExportFieldToAbsolutePath(exportField string) string {
	return filepath.Join(b.InstallLocation())
}

func (b *Plugin) CopyFileSheet(paths *target_paths.TargetPaths) (c *copy_file.CopyFile, hasFile bool) {
	s := b.Config().Get("styles").Str
	if s == "" {
		return &copy_file.CopyFile{}, false
	}
	return copy_file.NewCopyFileWithUniqueOutput(b.ExportFieldToAbsolutePath(s), func(uniqueId string) string {
		return filepath.Join(paths.GeneratedStyles(), fmt.Sprintf("%s.scss", uniqueId))
	}), true
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

func (p *Plugin) getSettingsPageData() plugin_setting_page_data.SettingsPageData {
	return plugin_setting_page_data.SettingsPageData{
		Title:      p.Config().Get("settings.title").Str,
		Subtitle:   p.Config().Get("settings.subtitle").Str,
		Href:       p.SettingsPageUrl(),
		PluginName: p.Name,
	}
}

func (p *Plugin) isValidSettingsData(data plugin_setting_page_data.SettingsPageData) bool {
	if (data.Href == "") || (data.PluginName == "") {
		return false
	}
	return true
}

func (p *Plugin) SettingsPageData() (data plugin_setting_page_data.SettingsPageData, ok bool) {
	item := p.getSettingsPageData()
	return item, p.isValidSettingsData(item)
}
