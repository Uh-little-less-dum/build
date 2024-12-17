package dependency

import (
	"fmt"

	"github.com/Uh-little-less-dum/build/pkg/types"
	"github.com/Uh-little-less-dum/build/pkg/utils"
)

type Dependency struct {
	name    string
	version string
	depType types.DependencyType
}

func (d *Dependency) Name() string {
	return d.name
}

func (d *Dependency) Version() string {
	return d.version
}

func (d *Dependency) Type() types.DependencyType {
	return d.depType
}

func (x *Dependency) InstallString() string {
	if x.version == "latest" {
		return x.name
	}
	return fmt.Sprintf("%s@%s", x.name, x.version)
}

func NewDependency(name, version string, depType ...types.DependencyType) Dependency {
	return Dependency{name: name, version: version, depType: utils.GetOptionalProperty(depType, types.ProductionDependency)}
}
