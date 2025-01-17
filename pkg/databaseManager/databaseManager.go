package database_manager

import (
	"embed"
	"fmt"
	"os/exec"
	"text/template"

	database_strategies "github.com/Uh-little-less-dum/build/pkg/databaseManager/stategies"
	database_types "github.com/Uh-little-less-dum/build/pkg/databaseManager/types"
	package_managers "github.com/Uh-little-less-dum/build/pkg/packageManager"
	target_paths "github.com/Uh-little-less-dum/build/pkg/targetPaths"
	"github.com/Uh-little-less-dum/go-utils/pkg/fs/file"
	"github.com/charmbracelet/log"
)

var (
	//go:embed "templates/*"
	templateFiles embed.FS
)

type DatabaseManager struct {
	paths    *target_paths.TargetPaths
	strategy database_strategies.DatabaseStrategy
	dbType   database_types.DatabaseType
}

func (d *DatabaseManager) SetDatabaseType(dbType database_types.DatabaseType) {
	d.dbType = dbType
}

func (d *DatabaseManager) SetRootPath(targetPaths *target_paths.TargetPaths) {
	d.paths = targetPaths
}

func (d *DatabaseManager) SetStrategy(s database_strategies.DatabaseStrategy) {
	d.strategy = s
}

func (d *DatabaseManager) generateNewDatabase(packageManager package_managers.PackageManager, targetDir string) *exec.Cmd {
	c := packageManager.RunScript("db:migrate")
	c.Dir = targetDir
	return c
}

func (d *DatabaseManager) migrateExistingDatabase(packageManager package_managers.PackageManager, targetDir string) *exec.Cmd {
	c := packageManager.RunScript("db:generate")
	c.Dir = targetDir
	return c
}

func (d *DatabaseManager) GenerateFromStrategy(packageManager package_managers.PackageManager, targetDir string) *exec.Cmd {
	switch d.strategy {
	case database_strategies.FromExisting:
		return d.migrateExistingDatabase(packageManager, targetDir)
	case database_strategies.GenerateNew:
		return d.generateNewDatabase(packageManager, targetDir)
	default:
		return d.generateNewDatabase(packageManager, targetDir)
	}
}

// absOutputPath must be absolute. Use the target_paths struct for reliability and maintainability, but pass just the schema path in here to avoid circular deps.
func (d *DatabaseManager) WriteSchemaFile(absOutputPath string) {
	templName := fmt.Sprintf("schema_%s.goprisma", d.dbType)
	templ, err := template.ParseFS(templateFiles, fmt.Sprintf("templates/%s", templName))
	if err != nil {
		log.Fatal(err)
	}
	f := file.NewFileItem(absOutputPath)
	err = templ.Execute(f, "")
	if err != nil {
		log.Fatal(err)
	}
}

func NewDatabaseManager(paths *target_paths.TargetPaths) *DatabaseManager {
	return &DatabaseManager{paths: paths, strategy: database_strategies.GenerateNew, dbType: database_types.Postgres}
}
