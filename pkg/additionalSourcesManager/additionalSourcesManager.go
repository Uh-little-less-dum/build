package additional_sources_manager

import (
	"fmt"
	target_paths "github.com/Uh-little-less-dum/build/pkg/targetPaths"
	copy_file "github.com/Uh-little-less-dum/go-utils/pkg/fs/copyFile"
	"github.com/Uh-little-less-dum/go-utils/pkg/fs/directory"
	"github.com/charmbracelet/log"
	"io/fs"
	"os"
	"path/filepath"
)

type AdditionalSourcesManager struct {
}

func (a *AdditionalSourcesManager) StylePaths(paths *target_paths.TargetPaths) []*copy_file.CopyFile {
	var res []*copy_file.CopyFile
	p, ok := a.Dir()
	if ok {
		stylesDir := filepath.Join(p, "styles")
		d := directory.NewDirectory(stylesDir)
		d.Walk(func(path string, d fs.DirEntry, err error) error {
			if !d.IsDir() {
				res = append(res, copy_file.NewCopyFileWithUniqueOutput(
					path,
					func(uniqueId string) string {
						return filepath.Join(paths.GeneratedStyles(), fmt.Sprintf("%s.scss", uniqueId))
					},
				))
			}
			return nil
		})
	}
	return res
}

func (a *AdditionalSourcesManager) ClearGeneratedStyles(paths *target_paths.TargetPaths) {
	dir := paths.GeneratedStyles()
	if dir != "" {
		d := directory.NewDirectory(dir)
		if d.Exists() {
			err := d.RemoveContents()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	// err := os.Mkdir(dir, 0777)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func (a *AdditionalSourcesManager) Dir() (dirPath string, ok bool) {
	p := os.Getenv("ULLD_ADDITIONAL_SOURCES")
	return p, p != ""
}

func NewAdditionalSourcesManager() *AdditionalSourcesManager {
	return &AdditionalSourcesManager{}
}
