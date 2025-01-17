package package_managers

import (
	"os/exec"

	"github.com/Uh-little-less-dum/build/pkg/types"
)

type Yarn struct {
	cwd string
}

func (n *Yarn) Id() PackageManagerId {
	return YarnId
}

func (n *Yarn) Key() string {
	return "yarn"
}

func (n *Yarn) Install() *exec.Cmd {
	return exec.Command("yarn", "install")
}

func (n *Yarn) Add(items []types.Installable) *exec.Cmd {
	k := []string{"add"}
	for _, l := range items {
		k = append(k, l.InstallString())
	}
	c := exec.Command("yarn", k...)
	c.Dir = n.cwd
	return c
}

func (n *Yarn) SetWorkingDir(workingDir string) {
	n.cwd = workingDir
}

func (n *Yarn) RunScript(cmds ...string) *exec.Cmd {
	c := append([]string{"run"}, cmds...)
	return exec.Command("yarn", c...)
}

func (n *Yarn) ModifiesPackageJson() bool {
	return false
}

func (n *Yarn) ModifyPackageJson(data []byte) []byte {
	return data
}
