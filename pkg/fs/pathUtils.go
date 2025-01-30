package fs_utils

import "strings"

func SplitPath(p string) []string {
	return strings.Split(p, "/")
}

func SplitPathWithTargetDir(subPath string, targetDir string) []string {
	vals := []string{targetDir}
	vals = append(vals, SplitPath(subPath)...)
	return vals
}
