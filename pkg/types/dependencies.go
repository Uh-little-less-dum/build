package types

type Installable interface {
	// Returns the install string for a specific package.
	// Example: redux or react-redux@2.15.21
	InstallString() string
}

type DependencyType string

const (
	ProductionDependency DependencyType = "dependencies"
	DevDependency        DependencyType = "devDependencies"
	PeerDependency       DependencyType = "peerDependencies"
	OptionalDependency   DependencyType = "optionalDependencies"
)

type Dependency interface {
	Name() string
	Version() string
	Type() DependencyType
}

type Plugin interface {
	Installable
}
