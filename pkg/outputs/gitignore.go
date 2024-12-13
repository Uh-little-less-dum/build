package outputs

import (
	"path/filepath"
)

type OutputFile_GitIgnore struct {
}

func (o OutputFile_GitIgnore) GetContent() string {
	return `node_modules
src/database/generated
.next
.env**
src/databaseTestData.ts
src/src/databaseTestData.ts
`
}

func (o OutputFile_GitIgnore) GetPath(rootDir string) string {
	return filepath.Join(rootDir, ".gitignore")
}
