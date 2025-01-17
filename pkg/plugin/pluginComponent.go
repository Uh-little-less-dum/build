package ulld_plugin

import (
	"fmt"
	"sync"

	plugin_component_docs_data "github.com/Uh-little-less-dum/build/pkg/plugin/componentDocs"
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

// FIX: Need to actually write the docs file here once a template is in order.
// Accepts the export string and a boolean indicating whether or not this is the summarzed for full documentation, and writes output to the filesystem. exportString has already been validated before being passed ib and is guaranteed to not be an empty string.
func (p *PluginComponent) writeDocsOutputFile(exportString string, isFullDocs bool) {

}

func (p *PluginComponent) WriteDocsOutput(wg *sync.WaitGroup) {
	dp, dpFull := p.DocPaths()
	if dp != "" {
		wg.Add(1)
		go func() {
			defer wg.Done()
			p.writeDocsOutputFile(dp, false)
		}()
	}
	if dpFull != "" {
		wg.Add(1)
		go func() {
			defer wg.Done()
			p.writeDocsOutputFile(dpFull, false)
		}()
	}
}

// FIX: Complete unimplemented... obviously.
func (p *PluginComponent) WriteSlotOutput(wg *sync.WaitGroup) {

}

func (p *PluginComponent) isValidComponentDocsData(data plugin_component_docs_data.ComponenDocData) bool {
	if (data.PluginName == "") || (data.FilePaths.Full == "" && (data.FilePaths.Short == "")) || (data.Urls.Full == "" && data.Urls.Short == "") {
		return false
	}
	return true
}

func (p *PluginComponent) getComponentDocDataItem(pluginName string) plugin_component_docs_data.ComponenDocData {
	tagData := p.data.Get("tags").Array()
	var tags = make([]string, len(tagData))
	for _, t := range tagData {
		tags = append(tags, t.Str)
	}
	// RESUME: Pick back up here.
	return plugin_component_docs_data.ComponenDocData{
		PluginName: pluginName,
	}
}

func (p *PluginComponent) GetComponentDocData(pluginName string) (data plugin_component_docs_data.ComponenDocData, ok bool) {
	item := p.getComponentDocDataItem(pluginName)
	return item, p.isValidComponentDocsData(item)
}

func NewPluginComponent(data gjson.Result, parentSlot slot_map.SlotKey) *PluginComponent {
	return &PluginComponent{data: data, parentSlot: parentSlot}
}
