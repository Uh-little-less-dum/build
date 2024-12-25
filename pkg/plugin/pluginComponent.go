package ulld_plugin

import (
	"fmt"

	slot_map "github.com/Uh-little-less-dum/build/pkg/slotMap"
	"github.com/tidwall/gjson"
)

type PluginComponent struct {
	// The json equivalent to components.# of the plugin's config.
	data       gjson.Result
	parentSlot slot_map.SlotKey
}

func (p *PluginComponent) ComponentName() string {
	return p.data.Get("componentName").Str
}

func (p *PluginComponent) ExportPath() string {
	return p.data.Get("export").Str
}

func (p *PluginComponent) Embeddables(pluginName, componentName, exportPath string, idx int) []PluginEmbeddableTemplateStruct {
	var res []PluginEmbeddableTemplateStruct
	embeddables := p.data.Get("embeddable").Array()
	for idx2, item := range embeddables {
		l := NewPluginEmbeddable(item).TemplateStruct(pluginName, componentName, exportPath, idx, idx2)
		res = append(res, l)
	}
	return res
}

func (p *PluginComponent) SlotKey() string {
	return fmt.Sprintf("%s/%s", p.parentSlot, p.data.Get("slot").Str)
}

func (p *PluginComponent) DocPaths() (shortDocs string, fullDocs string) {
	return p.data.Get("docsExport").Str, p.data.Get("fullDocsExport").Str
}

// Accepts the export string and a boolean indicating whether or not this is the summarzed for full documentation, and writes output to the filesystem. exportString has already been validated before being passed ib and is guaranteed to not be an empty string.
func (p *PluginComponent) WriteDocsOutput(exportString string, isFullDocs bool) {

}

func NewPluginComponent(data gjson.Result, parentSlot slot_map.SlotKey, pluginName string) *PluginComponent {
	return &PluginComponent{data: data, parentSlot: parentSlot}
}
