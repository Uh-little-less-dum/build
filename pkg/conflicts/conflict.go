package conflicts

type Conflict interface {
	Options() []string
	OnAccept(acceptedVal string)
}
