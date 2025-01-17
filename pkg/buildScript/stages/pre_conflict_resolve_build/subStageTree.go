package stage_pre_conflict_resolve_build

import (
	stage_write_npmrc "github.com/Uh-little-less-dum/build/pkg/buildScript/stages/pre_conflict_resolve_build/writeNpmrc"
	"github.com/Uh-little-less-dum/build/pkg/sub_stage"
	sub_command_ids "github.com/Uh-little-less-dum/go-utils/pkg/constants/subCommandIds"
	tea "github.com/charmbracelet/bubbletea"
)

func GetSubStageTree(program *tea.Program) []*sub_stage.SubStage {
	return []*sub_stage.SubStage{
		// Gather appConfig data
		sub_stage.NewSubStage("Gather App Config", "Taking a look at your config", sub_command_ids.GatherAppConfig, GatherAppConfig, program),
		// Gather root package.json data
		sub_stage.NewSubStage("Gather root package.json", "Grabbing some data from your package.json file", sub_command_ids.GatherRootPackageJson, GatherRootPackageJson, program),
		// Apply initial modifications to the root package.json file. Most of these are required to be able to keep the template app in the monorepo without installation conflicts.
		sub_stage.NewSubStage("Revise package.json", "Making some revisions to your package.json file", sub_command_ids.SetInitialPackageJsonData, SetInitialPackageJsonData, program),
		// Gather plugins from the package.json and the user's appConfig.
		sub_stage.NewSubStage("Gather Plugins", "Gathering your plugins", sub_command_ids.GatherPlugins, GatherPlugins, program),
		// TODO: Need to enable other package managers besides pnpm here.
		// Write the npmrc file for the user's selected package manager.
		sub_stage.NewSubStage("Writing package manager configuration", "Writing package manager configuration", sub_command_ids.WriteNpmrc, stage_write_npmrc.WriteNpmrc, program),
		// Allow each packageManager struct to modify the package.json file if required. Mostly useful for corepack.
		sub_stage.NewSubStage("Apply package manager specific settings", "Setting a few things to help your package manager", sub_command_ids.PackageManagerModifyPackageJson, PackageManagerModifyPackageJson, program),
		// Write modified package.json file to disk.
		sub_stage.NewSubStage("Save modified package.json", "Saving some configuration files", sub_command_ids.WritePackageJsonIfModified, WritePackageJsonIfModified, program),
		// Install dependencies from the newly modified package.json file.
		sub_stage.NewSubStage("Install Dependencies", "Installing dependencies. This might take a bit.", sub_command_ids.InstallDependencies, InstallDependencies, program),
		// Gather plugin conflicts before passing conflicts to user to resolve.
		sub_stage.NewSubStage("Gather plugin conflicts", "Checking for any conflicts amongst your plugins", sub_command_ids.GatherPluginConflicts, GatherPluginConflicts, program),
	}
}
