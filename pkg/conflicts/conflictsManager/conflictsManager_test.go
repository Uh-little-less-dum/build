package conflicts_handler

import (
	"testing"

	build_test "github.com/Uh-little-less-dum/build/internal/testing"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var testT *testing.T

// RESUME: Return here and finish this once the slot map struct is being automatically generated.
// FIX: This will pass, but that's just a coinki... this still needs to be implemented
func Test_GatherPluginConflicts(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	testT = t
	ginkgo.RunSpecs(t, "GatherPluginConflicts")
}

var _ = ginkgo.Describe("GatherPluginConflicts", func() {
	var mockManager *BuildConflictsManager
	ginkgo.BeforeEach(func() {
		mockManager = NewConflictsManager()
	})
	ginkgo.AfterEach(func() {
		mockManager = nil
	})
	var _ = ginkgo.Describe("", func() {
		ginkgo.When("No conflicts are provided", func() {
			ginkgo.It("Should have lengths of 0 for slot and page fields.", func() {
				mockManager.GatherPluginConflicts(build_test.GetMockPlugins(false, false))
				gomega.Expect(mockManager.Slot()).To(gomega.BeEmpty())
				gomega.Expect(mockManager.Page()).To(gomega.BeEmpty())
			})
		})
	})
})
