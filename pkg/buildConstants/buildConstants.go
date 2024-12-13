package build_constants

import "github.com/go-git/go-git/v5/plumbing"

type BuildConstant string

const (
	// Monorepo url.
	SparseCloneRepoUrl BuildConstant = "https://github.com/Uh-little-less-dum/ulld.git"
	// Path to sparse clone relative to monorepo.
	SparseCloneSparsePath BuildConstant = "apps/template"
	// Template repo url.
	TemplateRepoUrl    BuildConstant          = "https://github.com/Uh-little-less-dum/template.git"
	TemplateRemoteName string                 = "template"
	TemplateBranch     plumbing.ReferenceName = "template"
)
