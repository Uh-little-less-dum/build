package conflicts_page

type Conflict struct {
	Url string
}

func (c Conflict) Id() string {
	return c.Url
}

func (c Conflict) Options() []string {
	var res []string
	return res
}

func (c Conflict) OnAccept(acceptedVal string) {

}

func NewPageConflict(targetUrl string) *Conflict {
	return &Conflict{Url: targetUrl}
}
