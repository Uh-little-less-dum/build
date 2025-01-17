package event_handlers

import (
	"embed"
	"text/template"

	event_handler_types "github.com/Uh-little-less-dum/build/pkg/buildConstants/eventTypes"
	target_paths "github.com/Uh-little-less-dum/build/pkg/targetPaths"
	"github.com/Uh-little-less-dum/go-utils/pkg/fs/file"
	"github.com/charmbracelet/log"
	"github.com/tidwall/gjson"
)

var (
	//go:embed "templates/*"
	templateFiles embed.FS
)

type EventHandlerOfTypeList struct {
	eventType event_handler_types.PluginEventType
	items     []PluginEventHandler
}

func (e EventHandlerOfTypeList) EventType() event_handler_types.PluginEventType {
	return e.eventType
}

func (e EventHandlerOfTypeList) Items() []PluginEventHandler {
	return e.items
}

func (e EventHandlerOfTypeList) WriteOutput(paths *target_paths.TargetPaths) {
	templ, err := template.ParseFS(templateFiles, "templates/eventHandlerList.gotsx")
	if err != nil {
		log.Fatal(err)
	}
	outputPath := paths.EventMethodListOfType(e.eventType)
	f := file.NewFileItem(outputPath)
	f.WriteTemplate(templ, e)
}

// item: events.<event type> object
func (e *EventHandlerOfTypeList) Append(pluginName string, item gjson.Result) {
	e.items = append(e.items, NewPluginEventHandler(pluginName, item, len(e.items)))
}

func NewEventHandlerOfTypeList(eventType event_handler_types.PluginEventType) *EventHandlerOfTypeList {
	return &EventHandlerOfTypeList{
		eventType: eventType,
		items:     []PluginEventHandler{},
	}
}
