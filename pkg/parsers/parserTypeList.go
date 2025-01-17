package parsers

import ulld_plugin "github.com/Uh-little-less-dum/build/pkg/plugin"

type ParserOfTypeList struct {
	// parserKey: key of parser object in plugin config.
	parserKey string
	items     []ulld_plugin.PluginParser
}

func (p ParserOfTypeList) Items() []ulld_plugin.PluginParser {
	return p.items
}

func (p ParserOfTypeList) ParserType() string {
	return p.parserKey
}

func (p ParserOfTypeList) HasItems() bool {
	return len(p.items) > 0
}

func (p *ParserOfTypeList) Append(item ulld_plugin.PluginParser) {
	p.items = append(p.items, item)
}

func NewParserOfTypeList(parserKey string) *ParserOfTypeList {
	return &ParserOfTypeList{parserKey: parserKey}
}
