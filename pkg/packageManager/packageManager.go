package package_managers

import (
	"os/exec"

	"github.com/Uh-little-less-dum/build/pkg/types"
	"github.com/Uh-little-less-dum/build/pkg/utils"
	"github.com/elliotchance/orderedmap/v2"
)

type PackageManagerId int

const (
	NpmId PackageManagerId = iota
	PnpmId
	YarnId
	NoPackagekManagerSelect
)

func GetPackageManagerTitles() *orderedmap.OrderedMap[PackageManagerId, string] {
	// d := make(map[PackageManagerId]string)
	d := orderedmap.NewOrderedMap[PackageManagerId, string]()
	d.Set(PnpmId, "Pnpm")
	d.Set(NpmId, "Npm")
	d.Set(YarnId, "Yarn")
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
	// Returns a string representation of the package manager to be used in template strings.
	Key() string
	Install() *exec.Cmd
	Add(items []types.Installable) *exec.Cmd
	SetWorkingDir(workingDir string)
	RunScript(additionalCmds ...string) *exec.Cmd
	ModifyPackageJson(data []byte) []byte
	ModifiesPackageJson() bool
}

// WARN: Should alert user that they haven't provided a package manager and return the NoPackagekManagerSelected Id instead of defaulting to pnpm.
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
