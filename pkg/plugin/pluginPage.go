package ulld_plugin

import "github.com/tidwall/gjson"

type PluginPage struct {
	// The json equvialent of pages.#
	data gjson.Result
}

func (p *PluginPage) WriteOutput() {

}

func NewPluginPage(data gjson.Result) *PluginPage {
	return &PluginPage{data: data}
}
