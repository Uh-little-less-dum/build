package ulld_plugin

import (
	slot_map "github.com/Uh-little-less-dum/build/pkg/slotMap"
	target_paths "github.com/Uh-little-less-dum/build/pkg/targetPaths"
	"github.com/tidwall/gjson"
)

func NewPlugin(name, version string, slotKey slot_map.SlotKey, paths *target_paths.TargetPaths) *Plugin {
	return &Plugin{Name: name, Version: version, Slot: slotKey, paths: paths}
}

func NewPluginWithInstallLocation(name, version string, slotKey slot_map.SlotKey, paths *target_paths.TargetPaths, installLocation string) *Plugin {
	return &Plugin{Name: name, Version: version, Slot: slotKey, paths: paths, installLocation: installLocation}
}

// Takes the gjson.Result of a plugins.# item and returns the matching struct.
func PluginJsonToStruct(j gjson.Result, paths *target_paths.TargetPaths) *Plugin {
	return NewPlugin(j.Get("name").Str, j.Get("version").Str, slot_map.NoSlotApplied, paths)
}

func SlotJsonToStruct(k, v gjson.Result, paths *target_paths.TargetPaths) []*Plugin {
	res := []*Plugin{}
	items := v.Array()
	for _, l := range items {
		res = append(res, NewPlugin(l.Get("name").Str, l.Get("version").Str, slot_map.StringToSlotId(k.Str), paths))
	}
	return res
}
