package outputs

type OutputFile interface {
	// Gets content that should be written to file.
	GetContent() string
	// Gets absolute path of output file from provided targetDir.
	GetPath(targetDir string) string
}

// WARN: Will need to gather config specific files here as well.
func GetAllOutputFiles() []OutputFile {
	outputFiles := []OutputFile{
		OutputFile_Readme{},
		OutputFile_Npmrc{},
		OutputFile_GitIgnore{},
	}
	return outputFiles
}
