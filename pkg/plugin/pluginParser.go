package ulld_plugin

import (
	"fmt"

	parser_types "github.com/Uh-little-less-dum/go-utils/pkg/constants/parserTypes"
	"github.com/charmbracelet/log"
	"github.com/tidwall/gjson"
)

type PluginParser struct {
	// Json equivalent of plugins.#.parsers.#
	data gjson.Result
	// Key of parsers object denoting the type of parser (mdx, csv, etc...)
	parserKey string
}

type ParserItem struct {
	ParserType parser_types.ParserType
	ImportName string
	ImportPath string
}

func (p *PluginParser) ToTemplateData(idx, idx2 int) ParserItem {
	importPath := p.data.Get("export").Str
	if importPath == "" {
		log.Fatal("Attempted to get an empty importPath. Cannot continue.")
	}
	return ParserItem{ParserType: parser_types.ParserType(p.parserKey), ImportName: fmt.Sprintf("parser_%d%d", idx, idx2), ImportPath: importPath}
}

func NewPluginParser(k, v gjson.Result) *PluginParser {
	return &PluginParser{data: v, parserKey: k.Str}
}
