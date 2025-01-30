package ulld_plugin

import (
	"embed"
	"path/filepath"
	"text/template"

	target_paths "github.com/Uh-little-less-dum/build/pkg/targetPaths"
	"github.com/Uh-little-less-dum/go-utils/pkg/fs/file"
	"github.com/charmbracelet/log"
	"github.com/tidwall/gjson"
)

var (
	//go:embed "templates/*"
	templateFiles embed.FS
)

type PluginPage struct {
	// The json equvialent of pages.#
	data gjson.Result
}

type templateStruct struct{}

func (p *PluginPage) TargetUrl() string {
	return p.data.Get("targetUrl").Str
}

// FIX: This will never write page slots. Need to find more reliable way to check if a page is from a slot or not. Currently checking urls seems inadequate.
func (p *PluginPage) IsSlotPage(slotMapData gjson.Result) bool {
	return false
}

// FIX: Completely unimplemented. Need to check if a page is a slot page or not and write output properly. Should just move all slots to the seperate directory and embed them in the template app instead of trying to modify page.tsx files directly.
func (p *PluginPage) WriteSlotOutput(slotMapData gjson.Result, pluginName string, paths target_paths.TargetPaths) {

}

func (p *PluginPage) TemplateStruct() templateStruct {
	return templateStruct{}
}

func (p *PluginPage) OutputDir(paths *target_paths.TargetPaths) string {
	return paths.TargetUrlToDirname(p.TargetUrl())
}

func (p *PluginPage) WriteOutput(paths *target_paths.TargetPaths) {
	outputPath := filepath.Join(p.OutputDir(paths), "page.tsx")
	f := file.NewFileItem(outputPath)
	templ, err := template.ParseFS(templateFiles, "templates/pluginPage.gotsx")
	if err != nil {
		log.Fatal(err)
	}
	err = templ.ExecuteTemplate(f, "pluginPage.gotsx", p.TemplateStruct())
	if err != nil {
		log.Fatal(err)
	}
}

func NewPluginPage(data gjson.Result) *PluginPage {
	return &PluginPage{data: data}
}
