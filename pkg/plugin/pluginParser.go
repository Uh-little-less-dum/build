package ulld_plugin

import (
	"fmt"

	"github.com/tidwall/gjson"
)

type PluginParser struct {
	// Json equivalent of plugins.#.parsers.#
	data gjson.Result
	// Key of parsers object denoting the type of parser (mdx, csv, etc...)
	parserKey  string
	pluginName string
	idx        int
}

func (p PluginParser) ImportName() string {
	return fmt.Sprintf("%sParser_%d", p.parserKey, p.idx)
}

func (p PluginParser) ExportPath() string {
	return p.data.Get("export").Str
}

func (p PluginParser) ImportPath() string {
	exportPath := p.ExportPath()
	return fmt.Sprintf("%s/%s", p.pluginName, exportPath)
}

func (p PluginParser) ParserType() string {
	return p.parserKey
}

func NewPluginParser(k, v gjson.Result, pluginName string, idx int) *PluginParser {
	return &PluginParser{data: v, parserKey: k.Str, pluginName: pluginName, idx: idx}
}
