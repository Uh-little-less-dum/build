package package_managers

import (
	"os/exec"

	"github.com/Uh-little-less-dum/build/pkg/types"
)

type Pnpm struct {
	cwd string
}

func (n *Pnpm) Id() PackageManagerId {
	return PnpmId
}

func (npm *Pnpm) Install() *exec.Cmd {
	return exec.Command("pnpm", "install")
}

func (npm *Pnpm) Add(items []types.Installable) *exec.Cmd {
	k := []string{"add"}
	for _, l := range items {
		k = append(k, l.InstallString())
	}
	c := exec.Command("pnpm", k...)
	c.Dir = npm.cwd
	return c
}

func (n *Pnpm) SetWorkingDir(workingDir string) {
	n.cwd = workingDir
}
