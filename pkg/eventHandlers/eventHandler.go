package event_handlers

import (
	"fmt"

	event_handler_types "github.com/Uh-little-less-dum/build/pkg/buildConstants/eventTypes"
	"github.com/tidwall/gjson"
)

type PluginEventHandler struct {
	eventType  event_handler_types.PluginEventType
	pluginName string
	// exportPath: events.<event type> object in pluginConfig schema
	exportPath gjson.Result
	idx        int
}

func (p PluginEventHandler) ImportName() string {
	return fmt.Sprintf("%sEventHandler_%d", p.eventType, p.idx)
}

func (p PluginEventHandler) ImportPath() string {
	return fmt.Sprintf("%s/%s", p.PluginName(), p.exportPath.Str)
}

func (p PluginEventHandler) PluginName() string {
	return p.pluginName
}

func NewPluginEventHandler(pluginName string, data gjson.Result, idx int) PluginEventHandler {
	return PluginEventHandler{pluginName: pluginName, exportPath: data, idx: idx}
}
