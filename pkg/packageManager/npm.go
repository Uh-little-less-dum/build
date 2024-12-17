package package_managers

import (
	"os/exec"

	"github.com/Uh-little-less-dum/build/pkg/types"
)

type Npm struct {
	cwd string
}

func (n *Npm) Id() PackageManagerId {
	return NpmId
}

func (n *Npm) Install() *exec.Cmd {
	return exec.Command("npm", "install")
}

func (n *Npm) Add(items []types.Installable) *exec.Cmd {
	k := []string{"install"}
	for _, l := range items {
		k = append(k, l.InstallString())
	}
	c := exec.Command("npm", k...)
	c.Dir = n.cwd
	return c
}

func (n *Npm) SetWorkingDir(workingDir string) {
	n.cwd = workingDir
}
