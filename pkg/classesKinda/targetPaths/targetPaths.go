package target_paths

import (
	"path/filepath"

	"github.com/charmbracelet/log"
)

type TargetPaths struct {
	projectRoot string
}

// Path to the root of the public directory.
func (t TargetPaths) Public() string {
	return filepath.Join(t.projectRoot, "public")
}

// Path to styles directory containing all css and scss files.
func (t TargetPaths) Styles() string {
	return filepath.Join(t.projectRoot, "src", "styles")
}

// Path to package.json file
func (t TargetPaths) PackageJson() string {
	return filepath.Join(t.projectRoot, "package.json")
}

// Path to the app directory at src/app
func (t TargetPaths) App() string {
	return filepath.Join(t.projectRoot, "src", "app")
}
func (t TargetPaths) ProjectRoot() string {
	return t.projectRoot
}

// Path to the *directory* containing all local, automatically gathered documentation.
func (t TargetPaths) LocalDocumentation() string {
	return filepath.Join(t.projectRoot, "generatedMarkdown")
}
func (t TargetPaths) PluginDocumentation() string {
	return filepath.Join(t.LocalDocumentation(), "pluginDocs")
}
func (t TargetPaths) ComponentDocumentation() string {
	return filepath.Join(t.LocalDocumentation(), "componentDocs")
}
func (t TargetPaths) FullComponentDocumentation() string {
	return filepath.Join(t.ComponentDocumentation(), "full")
}
func (t TargetPaths) ShortComponentDocumentation() string {
	return filepath.Join(t.ComponentDocumentation(), "short")
}

// Path to node_modules directory.
func (t TargetPaths) Node_modules() string {
	return filepath.Join(t.projectRoot, "node_modules")
}

// Path to src directory
func (t TargetPaths) Src() string {
	return filepath.Join(t.projectRoot, "src")
}

// Path to app directory at src/app
func (t TargetPaths) AppDir() string {
	return filepath.Join(t.projectRoot, "src", "app")
}

// Path to next.js config file.
func (t TargetPaths) NextConfig() string {
	return filepath.Join(t.projectRoot, "next.config.mjs")
}

// Path to tailwind.config.ts
func (t TargetPaths) Tailwind() string {
	return filepath.Join(t.projectRoot, "tailwind.config.ts")
}

// Path to ulldBuildData.json file.
func (t TargetPaths) UlldBuildData() string {
	return filepath.Join(t.projectRoot, "ulldBuildData.json")
}

// Path to the main appConfig file that's used during compilation and runtime.
func (t TargetPaths) AppConfig() string {
	return filepath.Join(t.projectRoot, "appConfig.ulld.json")
}

// Path to the unified component map file.
func (t TargetPaths) ComponentMap() string {
	return filepath.Join(t.projectRoot, "src", "internal", "componentMap.ts")
}
func (t TargetPaths) OnBackupMethodList() string {
	return filepath.Join(t.projectRoot, "events", "methodLsts", "backupMethods.ts")
}
func (t TargetPaths) OnRestoreMethodList() string {
	return filepath.Join(t.projectRoot, "events", "methodLists", "restoreMethods.ts")
}
func (t TargetPaths) OnSyncMethodList() string {
	return filepath.Join(t.projectRoot, "events", "methodLists", "syncMethods.ts")
}
func (t TargetPaths) OnBuildMethodList() string {
	return filepath.Join(t.projectRoot, "buildUtils", "__TEMP__", "onBuildMethodList.ts")
}

// Path to the src/methods directory that contains all event related methods (onSync, onBackup, etc...)
func (t TargetPaths) Methods() string {
	return filepath.Join(t.projectRoot, "src", "methods")
}
func (t TargetPaths) MdxParserList() string {
	return filepath.Join(t.projectRoot, "src", "methods", "parsers", "parserLists", "mdx.ts")
}

// Returns the temporary *directory* path within the build output.
func (t TargetPaths) TempBuildFiles() string {
	return filepath.Join(t.projectRoot, "buildUtils", "__TEMP__")
}

// Path to the root of the directory where css files should be gathered.
func (t TargetPaths) UserDefinedStyles() string {
	return filepath.Join(t.Styles(), "userProvided")
}

// Path to the schema.prisma file
func (t TargetPaths) PrismaSchema() string {
	return filepath.Join(t.projectRoot, "src", "database", "schema.prisma")
}

// .gitignore file
func (t TargetPaths) Gitignore() string {
	return filepath.Join(t.projectRoot, ".gitignore")
}

// .env.local file
func (t TargetPaths) EnvLocal() string {
	return filepath.Join(t.projectRoot, ".env.local")
}

func (t TargetPaths) XdgPaths() string {
	log.Fatal("Need to enable the xdg paths here when back on wifi and able to install it again")
	return ""
}

func NewTargetPaths(rootDir string) TargetPaths {
	return TargetPaths{projectRoot: rootDir}
}
