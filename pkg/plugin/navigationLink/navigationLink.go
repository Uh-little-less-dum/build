package navigation_link

import "github.com/tidwall/gjson"

type NavigationLink struct {
	Label      string `json:"label,omitempty"`
	Href       string `json:"href,omitempty"`
	Icon       string `json:"icon,omitempty"`
	PluginName string `json:"pluginName,omitempty"`
	Category   string `json:"category,omitempty"`
}

func NewNavigationLink(data gjson.Result) NavigationLink {
	return NavigationLink{
		Label:      data.Get("label").Str,
		Href:       data.Get("href").Str,
		Icon:       data.Get("icon").Str,
		PluginName: data.Get("pluginName").Str,
		Category:   data.Get("category").Str,
	}
}
