package ulld_plugin_test

import (
	"path/filepath"
	"testing"

	"github.com/Uh-little-less-dum/build/mocks"
	ulld_plugin "github.com/Uh-little-less-dum/build/pkg/plugin"
	slot_map "github.com/Uh-little-less-dum/build/pkg/slotMap"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var testT *testing.T

func Test_GatherPluginConflicts(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	testT = t
	ginkgo.RunSpecs(t, "GatherPluginConflicts")
}

// TODO: Add an output to the gatherCliData method that returns only plugins with valid configs, and map over each plugin here. That should catch some errors on the typescript side.
// RESUME: Come back here and implement the rest of these. Today didn't feel very productive, but all of the boilerplate is pretty much handled.
var _ = ginkgo.Describe("Plugin suite", func() {
	var plugin *ulld_plugin.Plugin
	ginkgo.BeforeEach(func() {
		plugin = ulld_plugin.NewPlugin("@ulld/navigation", "1.0.0", slot_map.Bibliography, mocks.TargetPaths())
	})
	ginkgo.AfterEach(func() {
		plugin = nil
	})
	var _ = ginkgo.Describe("InstallLocation", func() {
		ginkgo.When("InstallLocation is called with valid installation", func() {
			ginkgo.It("Should return the directory to the root of the installed dependency", func() {
				il := plugin.InstallLocation()
				gomega.Expect(il).To(gomega.BeADirectory())
				gomega.Expect(filepath.Join(il, "package.json")).To(gomega.BeAnExistingFile())
				gomega.Expect(filepath.Join(il, "pluginConfig.ulld.json")).To(gomega.BeAnExistingFile())
			})
		})
	})
	var _ = ginkgo.Describe("Getting config", func() {
		ginkgo.When("Config is present", func() {
			ginkgo.It("returns a valid gjson struct", func() {
				name := plugin.Config().Get("pluginName").Str
				gomega.Expect(name).To(gomega.Equal(plugin.Name))
			})
		})
	})
	var _ = ginkgo.Describe("Get Components", func() {
		ginkgo.When("Config is present", func() {
			ginkgo.It("Returns valid component structs", func() {
				comps := plugin.Components()
				gomega.Expect(comps).NotTo(gomega.BeEmpty())
			})
		})
	})
})
