package cleanup_before_transpile_stages

import (
	"embed"

	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
	"github.com/Uh-little-less-dum/go-utils/pkg/fs/file"
	"github.com/charmbracelet/log"
)

var (
	//go:embed "templates/*"
	templateFiles embed.FS
)

func WriteGitIgnore(cfg *build_config.BuildManager) {
	p := cfg.Paths.Gitignore()
	f := file.NewFileItem(p)
	b, err := templateFiles.ReadFile("templates/gitignore.txt")
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write(b)
	if err != nil {
		log.Fatal(err)
	}
}
