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

func (n *Npm) Key() string {
	return "npm"
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

func (n *Npm) RunScript(cmds ...string) *exec.Cmd {
	c := append([]string{"run"}, cmds...)
	return exec.Command("npm", c...)
}

func (n *Npm) ModifiesPackageJson() bool {
	return false
}

func (n *Npm) ModifyPackageJson(data []byte) []byte {
	return data
}
