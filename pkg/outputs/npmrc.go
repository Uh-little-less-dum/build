package outputs

import (
	"path/filepath"
)

type OutputFile_Npmrc struct {
}

func (o OutputFile_Npmrc) GetContent() string {
	return `hoist-workspace-packages=true
public-hoist-pattern[]=*prisma*
auto-install-peers=true
node-linker=hoisted
//registry.npmjs.org/:_authToken=${NPM_TOKEN}`
}

func (o OutputFile_Npmrc) GetPath(rootDir string) string {
	return filepath.Join(rootDir, ".npmrc")
}
