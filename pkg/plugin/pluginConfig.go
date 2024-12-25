package ulld_plugin

import (
	"os"
	"path/filepath"

	fs_utils "github.com/Uh-little-less-dum/build/pkg/fs"
	"github.com/charmbracelet/log"
	"github.com/tidwall/gjson"
)

func logConfigFail(name string) {
	log.Fatalf("Attempted to gather the plugin config for %s unsuccessfully.", name)
}

type PluginConfig struct {
	hasReadConfig bool
	data          gjson.Result
}

func (p *PluginConfig) configFromPackageJson(installLoc, pluginName string) gjson.Result {
	fp := filepath.Join(installLoc, "package.json")
	if !fs_utils.Exists(fp) {
		logConfigFail(pluginName)
	}
	b, err := os.ReadFile(fp)
	if err != nil {
		logConfigFail(pluginName)
	}
	j := gjson.ParseBytes(b)
	l := j.Get("ulld-pluginConfig")
	if l.Str == "" {
		logConfigFail(pluginName)
	}
	p.hasReadConfig = true
	p.data = j
	return l
}

func (p *PluginConfig) gatherConfig(installLocation, pluginName string) gjson.Result {
	fp := filepath.Join(installLocation, "pluginConfig.ulld.json")
	if !fs_utils.Exists(fp) {
		return p.configFromPackageJson(installLocation, pluginName)
	}
	b, err := os.ReadFile(fp)
	if err != nil {
		log.Fatal(err)
	}
	j := gjson.ParseBytes(b)
	p.hasReadConfig = true
	p.data = j
	return j
}

func (p *PluginConfig) Config(installLocation, pluginName string) gjson.Result {
	if p.hasReadConfig {
		return p.data
	}
	return p.gatherConfig(installLocation, pluginName)
}
