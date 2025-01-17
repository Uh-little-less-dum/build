package stage_write_npmrc

import (
	"embed"
	"fmt"
	"path/filepath"
	"text/template"

	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
	"github.com/Uh-little-less-dum/go-utils/pkg/fs/file"
	"github.com/charmbracelet/log"
)

var (
	//go:embed "templates/*"
	templateFiles embed.FS
)

func WriteNpmrc(cfg *build_config.BuildManager) {
	targetPath := filepath.Join(cfg.TargetDir(), ".npmrc")
	templPage := fmt.Sprintf("%s.gotxt", cfg.PackageManager().Key())
	templ, err := template.ParseFS(templateFiles, filepath.Join("templates", templPage))
	if err != nil {
		log.Fatal(err)
	}
	f := file.NewFileItem(targetPath)
	fi := f.CreateIfNotExists()
	defer fi.Close()
	err = templ.ExecuteTemplate(fi, templPage, "")
	if err != nil {
		log.Fatal(err)
	}
}
