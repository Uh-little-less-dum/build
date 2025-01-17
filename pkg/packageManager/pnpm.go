package package_managers

import (
	"os/exec"

	"github.com/Uh-little-less-dum/build/pkg/types"
	"github.com/charmbracelet/log"
	"github.com/tidwall/sjson"
)

type Pnpm struct {
	cwd string
}

func (n *Pnpm) Id() PackageManagerId {
	return PnpmId
}

func (n *Pnpm) Key() string {
	return "pnpm"
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

func (n *Pnpm) RunScript(cmds ...string) *exec.Cmd {
	return exec.Command("pnpm", cmds...)
}

type peerDepsData struct {
	ignoreMissing []string
	allowAny      []string
}

type pnpmPackageJsonData struct {
	peerDependencyRules peerDepsData
}

func getModifiedPnpmPackageJsonData() pnpmPackageJsonData {
	return pnpmPackageJsonData{
		peerDependencyRules: peerDepsData{
			ignoreMissing: []string{
				"@babel/*",
				"@eslint/*",
				"**eslint**",
			},
			allowAny: []string{
				"@babel/*",
				"eslint",
				"**eslint**",
			},
		},
	}
}

func (n *Pnpm) ModifiesPackageJson() bool {
	return true
}

func (n *Pnpm) ModifyPackageJson(data []byte) []byte {
	newData, err := sjson.SetBytes(data, "pnpm", getModifiedPnpmPackageJsonData())
	if err != nil {
		log.Fatal(err)
	}
	return newData
}
