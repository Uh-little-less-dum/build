package styles_manager

import (
	"cmp"
	"fmt"
	"sync"

	target_paths "github.com/Uh-little-less-dum/build/pkg/targetPaths"
	copy_file "github.com/Uh-little-less-dum/go-utils/pkg/fs/copyFile"
	"github.com/Uh-little-less-dum/go-utils/pkg/fs/file"
	"github.com/charmbracelet/log"
)

type StylesManager struct {
}

func (m *StylesManager) WriteScssFiles(paths *target_paths.TargetPaths, items []*copy_file.CopyFile) {
	var wg sync.WaitGroup
	var outputPaths = make(chan string, len(items))
	for _, item := range items {
		wg.Add(1)
		go func() {
			defer wg.Done()
			outputPaths <- item.WriteOutput()
		}()
	}
	wg.Wait()
	close(outputPaths)
	tunnelPath := paths.GeneratedStylesTunnelFile()
	var s string
	for p := range outputPaths {
		s += fmt.Sprintf("@use \"%s.scss\";\n", paths.TunnelRelativeStyleSheet(p))
	}
	f := file.NewFileItem(tunnelPath)
	_, err := f.WriteForcefully([]byte(cmp.Or(s, "// No generated scss files.")))
	if err != nil {
		log.Fatal(err)
	}
}

func NewStylesManager() *StylesManager {
	return &StylesManager{}
}
