package conflicts_slot

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/log"
)

type Conflict struct {
	ParentSlot string
	SubSlot    string
	Resolved   bool
	// acceptedVal: pluginName
	AcceptedVal string
}

func (c Conflict) Id() string {
	return fmt.Sprintf("%s/%s", c.ParentSlot, c.SubSlot)
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

func NewSlotConflict(slotString string) *Conflict {
	s := strings.Split(slotString, "/")
	if len(s) != 2 {
		log.Fatalf("Found a slot conflict for the %s slot that seems to be due to a malformed config on the plugin's end. Cannot continue.", slotString)
	}
	return &Conflict{ParentSlot: s[0], SubSlot: s[1]}
}
