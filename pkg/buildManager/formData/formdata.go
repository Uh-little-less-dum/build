package form_data

type BuildFormData struct {
	targetDir string
}

func (b *BuildFormData) TargetDir() string {
	return b.targetDir
}

func (b *BuildFormData) SetTargetDirOnlyInBuildConfig(d string) {
	b.targetDir = d
}
