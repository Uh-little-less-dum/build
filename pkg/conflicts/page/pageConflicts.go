package conflicts_page

type Conflict struct {
	Url      string
	Resolved bool
	// acceptedVal: pluginName
	AcceptedVal string
}

func (c Conflict) Id() string {
	return c.Url
}

func (c Conflict) Options() []string {
	var res []string
	return res
}

// acceptedVal: pluginName
func (c *Conflict) OnAccept(acceptedVal string) {
	c.AcceptedVal = acceptedVal
	c.Resolved = true
}

func NewPageConflict(targetUrl string) *Conflict {
	return &Conflict{Url: targetUrl}
}
