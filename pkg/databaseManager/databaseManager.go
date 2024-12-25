package database_manager

import (
	"os/exec"

	database_strategies "github.com/Uh-little-less-dum/build/pkg/databaseManager/stategies"
	package_managers "github.com/Uh-little-less-dum/build/pkg/packageManager"
	target_paths "github.com/Uh-little-less-dum/build/pkg/targetPaths"
)

type DatabaseManager struct {
	paths    *target_paths.TargetPaths
	strategy database_strategies.DatabaseStrategy
}

func (d *DatabaseManager) SetRootPath(targetPaths *target_paths.TargetPaths) {
	d.paths = targetPaths
}

func (d *DatabaseManager) SetStrategy(s database_strategies.DatabaseStrategy) {
	d.strategy = s
}

// RESUME: Need to implement this with the proper scripts here. These are almost surely half-assed.
func (d *DatabaseManager) generateNewDatabase(packageManager package_managers.PackageManager) *exec.Cmd {
	return packageManager.RunScript("prisma", "migrate")
}

// RESUME: Need to implement this with the proper scripts here. These are almost surely half-assed.
func (d *DatabaseManager) migrateExistingDatabase(packageManager package_managers.PackageManager) *exec.Cmd {
	return packageManager.RunScript("prisma", "generate")
}

func (d *DatabaseManager) GenerateFromStrategy(packageManager package_managers.PackageManager) *exec.Cmd {
	switch d.strategy {
	case database_strategies.FromExisting:
		return d.migrateExistingDatabase(packageManager)
	case database_strategies.GenerateNew:
		return d.generateNewDatabase(packageManager)
	default:
		return d.generateNewDatabase(packageManager)
	}
}

func NewDatabaseManager(paths *target_paths.TargetPaths) *DatabaseManager {
	return &DatabaseManager{paths: paths, strategy: database_strategies.GenerateNew}
}
