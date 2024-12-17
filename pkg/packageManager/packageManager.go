package package_managers

import (
	"os/exec"

	"github.com/Uh-little-less-dum/build/pkg/types"
	"github.com/Uh-little-less-dum/build/pkg/utils"
)

type PackageManagerId int

const (
	NpmId PackageManagerId = iota
	PnpmId
	YarnId
	NoPackagekManagerSelected
)

func GetPackageManagerTitles() map[PackageManagerId]string {
	d := make(map[PackageManagerId]string)
	d[PnpmId] = "Pnpm"
	d[NpmId] = "Npm"
	d[YarnId] = "Yarn"
	return d
}

func GetAvailablePackageManagers() map[PackageManagerId]bool {
	d := make(map[PackageManagerId]bool)
	_, d[PnpmId] = utils.IsInstalled("pnpm")
	_, d[NpmId] = utils.IsInstalled("npm")
	_, d[YarnId] = utils.IsInstalled("yarn")
	return d
}

type PackageManager interface {
	Id() PackageManagerId
	Install() *exec.Cmd
	Add(items []types.Installable) *exec.Cmd
	SetWorkingDir(workingDir string)
}

// FIX: Should alert user that they haven't provided a package manager and return the NoPackagekManagerSelected Id instead of defaulting to pnpm.
func GetPackageManagerStruct(id PackageManagerId) PackageManager {
	var n PackageManager = &Pnpm{}
	switch id {
	case NpmId:
		n = &Npm{}
	case YarnId:
		n = &Yarn{}
	case PnpmId:
		n = &Pnpm{}
	}
	return n
}
