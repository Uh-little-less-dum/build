package ulld_plugin

import (
	"fmt"

	"github.com/Uh-little-less-dum/build/pkg/buildConstants/slots"
	"github.com/tidwall/gjson"
)

type Plugin struct {
	Name    string
	Version string
	Slot    slots.SlotKey
}

// Returns the install string for a specific package.
// Example: redux or react-redux@2.15.21
func (plugin *Plugin) InstallString() string {
	if plugin.Version == "latest" {
		return plugin.Name
	}
	return fmt.Sprintf("%s@%s", plugin.Name, plugin.Version)
}

// Takes the gjson.Result of a plugins.# item and returns the matching struct.
func PluginJsonToStruct(j gjson.Result) *Plugin {
	return &Plugin{Name: j.Get("name").Str, Version: j.Get("version").Str, Slot: slots.NoSlotApplied}
}

func SlotJsonToStruct(k, v gjson.Result) []*Plugin {
	res := []*Plugin{}
	items := v.Array()
	for _, l := range items {
		res = append(res, &Plugin{Name: l.Get("name").Str, Version: l.Get("version").Str, Slot: slots.StringToSlotId(k.Str)})
	}
	return res
}
