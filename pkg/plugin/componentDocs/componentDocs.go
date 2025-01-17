package plugin_component_docs_data

type componentDocsGroup struct {
	Short string `json:"short"`
	Full  string `json:"full"`
}

type ComponenDocData struct {
	PluginName       string             `json:"pluginName"`
	ComponentName    string             `json:"componentName"`
	EmbeddableSyntax []string           `json:"embeddableSyntax"`
	Urls             componentDocsGroup `json:"urls"`
	FilePaths        componentDocsGroup `json:"filePaths"`
	Tags             []string           `json:"tags"`
	ComponentId      string             `json:"componentId"`
}
