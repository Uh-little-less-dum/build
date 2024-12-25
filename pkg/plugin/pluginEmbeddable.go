package ulld_plugin

import (
	"fmt"

	"github.com/tidwall/gjson"
)

type PluginEmbeddable struct {
	// The json equivalent of components.#.embeddable
	data       gjson.Result
	pluginName *string
	// The export path relative to the plugins root.
	// Example: "embbeddableEquatioin" instead of "@ulld/equatioins/embeddableEquation"
	exportPath *string
}

func (p PluginEmbeddable) importPath() {
	panic("Not yet implemented")
}

func (p PluginEmbeddable) importName(componentName string, idx1, idx2 int) string {
	return fmt.Sprintf("%s_%d%d", componentName, idx1, idx2)
}

func (p PluginEmbeddable) label() string {
	return p.data.Get("label").Str
}

func (p PluginEmbeddable) regex() string {
	return p.data.Get("regexToInclude").Str
}

type PluginEmbeddableTemplateStruct struct {
	ImportName, Regex, Label, ImportPath string
}

func (p PluginEmbeddable) TemplateStruct(pluginName, componentName, exportPath string, idx1, idx2 int) PluginEmbeddableTemplateStruct {
	return PluginEmbeddableTemplateStruct{
		ImportName: p.importName(componentName, idx1, idx2),
		Label:      p.label(),
		Regex:      p.regex(),
		ImportPath: fmt.Sprintf("%s/%s", pluginName, exportPath),
	}
}

func NewPluginEmbeddable(data gjson.Result) *PluginEmbeddable {
	return &PluginEmbeddable{data: data}
}
