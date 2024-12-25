package ulld_plugin

import "github.com/tidwall/gjson"

type PluginSettings struct {
	// Represents the entire pluginConfig, not just the setting field.
	data *gjson.Result
}

func (p *PluginSettings) WriteOutput() {

}

func NewPluginSettings(data *gjson.Result) *PluginSettings {
	return &PluginSettings{data: data}
}
