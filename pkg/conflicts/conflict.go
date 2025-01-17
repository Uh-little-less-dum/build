package conflicts

type Conflict interface {
	Options() []string
	// acceptedVal: pluginName
	OnAccept(acceptedVal string)
}
