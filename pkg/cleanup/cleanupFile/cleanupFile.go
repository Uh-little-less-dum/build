package cleanup_file

import (
	"os"
	"path/filepath"

	"github.com/samber/lo"
)

// A file or directory that should be completely removed after the build is complete.
type CleanupFile struct {
	// Path relative to targetDir
	path string
}

func (c CleanupFile) Remove(targetDir string) {
	absPath := filepath.Join(targetDir, c.path)
	os.RemoveAll(absPath)
}

func GetCleanupFiles(relativePaths []string) []CleanupFile {
	return lo.Map[string, CleanupFile](relativePaths, func(item string, _ int) CleanupFile {
		return CleanupFile{item}
	})
}
