package post_conflict_resolve_stages

import (
	"embed"
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

// Writes onSync handlers as a single function to the target file.
func WriteOnSyncHandlers(cfg *build_config.BuildManager) {
	var handlers []ulld_plugin.PluginEventHandler
	for i, p := range cfg.Plugins {
		a, ok := p.Events().OnSyncMethod(i)
		if ok {
			handlers = append(handlers, a)
		}
	}

	templ, err := template.ParseFS(templateFiles, "templates/eventHandlerList.gotsx")
	if err != nil {
		log.Fatal(err)
	}
	f := file.NewFileItem(cfg.Paths.OnSyncMethodList())
	err = templ.ExecuteTemplate(f, "eventHandlerList.gotsx", handlers)
	if err != nil {
		log.Fatal(err)
	}
}
