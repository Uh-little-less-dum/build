package ulld_plugin

import (
	"fmt"
	"text/template"

	target_paths "github.com/Uh-little-less-dum/build/pkg/targetPaths"
	"github.com/Uh-little-less-dum/go-utils/pkg/fs/file"
	"github.com/charmbracelet/log"
	"github.com/tidwall/gjson"
)

type PluginSettings struct {
	// Represents the entire pluginConfig, not just the setting field.
	data *gjson.Result
}

func (p PluginSettings) PluginName() string {
	return p.data.Get("pluginName").Str
}

func (p PluginSettings) TargetUrl() string {
	return p.data.Get("settings.targetUrl").Str
}

func (p PluginSettings) ImportPath() string {
	return fmt.Sprintf("%s/%s", p.PluginName(), p.data.Get("settings.settingPageExport").Str)
}

func (p PluginSettings) HasSettingsPage() bool {
	return p.data.Get("settings.settingPageExport").Exists()
}

func (p PluginSettings) WriteOutput(paths *target_paths.TargetPaths) {
	if !p.HasSettingsPage() {
		return
	}
	templ, err := template.ParseFS(templateFiles, "templates/pluginSettingsPage.gotsx")
	if err != nil {
		log.Fatal(err)
	}
	f := file.NewFileItem(paths.TargetUrlToDirname(p.TargetUrl()))
	err = templ.ExecuteTemplate(f, "pluginSettingsPage.gotsx", p)
	if err != nil {
		log.Fatal(err)
	}
}

func NewPluginSettings(data *gjson.Result) *PluginSettings {
	return &PluginSettings{data: data}
}
