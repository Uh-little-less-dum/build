package parsers

import (
	"embed"
	"sync"
	"text/template"

	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
	ulld_plugin "github.com/Uh-little-less-dum/build/pkg/plugin"
	"github.com/Uh-little-less-dum/go-utils/pkg/fs/file"
	"github.com/charmbracelet/log"
)

var (
	//go:embed "templates/*"
	templateFiles embed.FS
)

type ParserList struct {
	items map[string]*ParserOfTypeList
}

func (p ParserList) AppendByType(item ulld_plugin.PluginParser) {
	pt := item.ParserType()
	p.items[pt].Append(item)
}

func (p *ParserList) WriteOutput(cfg *build_config.BuildManager) {
	var wg sync.WaitGroup
	for k, v := range p.items {
		// if v.HasItems() {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fp := cfg.Paths.ParserListOfType(k)
			f := file.NewFileItem(fp)
			templ, err := template.ParseFS(templateFiles, "templates/parserList.gotsx")
			if err != nil {
				log.Fatal(err)
			}
			err = templ.ExecuteTemplate(f, "parserList.gotsx", *v)
			if err != nil {
				log.Fatal(err)
			}
		}()
		// }
	}
	wg.Wait()
}

func newParserTypeMap() map[string]*ParserOfTypeList {
	m := make(map[string]*ParserOfTypeList)
	pts := ParserTypes()
	for _, pt := range pts {
		m[pt] = NewParserOfTypeList(pt)
	}
	return m
}

func NewParserLists(cfg *build_config.BuildManager) *ParserList {
	p := ParserList{
		items: newParserTypeMap(),
	}
	for _, l := range cfg.Plugins {
		for _, parser := range l.Parsers() {
			p.AppendByType(*parser)
		}
	}
	return &p
}
