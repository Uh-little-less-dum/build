package post_conflict_resolve_stages

import (
	"text/template"

	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
	"github.com/Uh-little-less-dum/go-utils/pkg/fs/file"
	"github.com/charmbracelet/log"
)

func WriteComponentMap(cfg *build_config.BuildManager) {
	data := cfg.Embeddables()
	tmpl, err := template.ParseFS(templateFiles, "templates/componentMap.gotsx")
	if err != nil {
		log.Fatal(err)
	}
	f := file.NewFileItem(cfg.Paths.ComponentMap())
	err = tmpl.ExecuteTemplate(f, "componentMap.gotsx", data)
	if err != nil {
		log.Fatal(err)
	}
}
