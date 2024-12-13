package outputs

import (
	"path/filepath"
)

type OutputFile_Readme struct {
}

func (o OutputFile_Readme) GetContent() string {
	return `# Uh Little Less Dum

Welcome to your new ULLD application. Make sure to visit the blog [here](https://uhlittlelessdum.com/blog) to stay up to date on the latest releases, future plans, and to useful guides.  

If you can, please consider visiting our sponsor page [here](https://uhlittlelessdum.com/sponsor). It's the middle of winter in Wisconsin, and I'm still quite literally homeless.

> This app was generated using the ULLD cli. Learn more by visiting [UhLittleLessDum.com](https://uhlittlelessdum.com)
`
}

func (o OutputFile_Readme) GetPath(rootDir string) string {
	return filepath.Join(rootDir, "README.md")
}
