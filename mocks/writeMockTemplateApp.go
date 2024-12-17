package mocks

import (
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/spf13/afero"
)

type MockFileType int

const (
	File MockFileType = iota
	Dir
)

type MockTemplateFile struct {
	TemplatePath string
	DataType     MockFileType
}

func getMockFiles() []MockTemplateFile {
	return []MockTemplateFile{

		{TemplatePath: "package.json", DataType: File},
		{TemplatePath: "next.config.mjs", DataType: File},
		{TemplatePath: "tailwind.config.ts", DataType: File},
	}
}

func WriteMockTemplateApp(f afero.Fs) {
	items := getMockFiles()
	buildRoot := os.Getenv("ULLD_BUILD_DEV_ROOT")
	if buildRoot == "" {
		log.Fatal("Could not find ULLD_BUILD_DEV_ROOT env variable.")
	}
	for _, item := range items {
		if item.DataType == File {
			content, err := os.ReadFile(filepath.Join(buildRoot, "mocks", "fileContent", item.TemplatePath))
			if err != nil {
				log.Fatal(err)
			}
			afero.WriteFile(f, item.TemplatePath, content, 0777)
		}
	}
}
