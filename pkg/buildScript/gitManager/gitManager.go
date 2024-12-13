package git_manager

import (
	"context"
	"os"
	"time"

	build_constants "github.com/Uh-little-less-dum/build/pkg/buildConstants"
	viper_keys "github.com/Uh-little-less-dum/go-utils/pkg/constants/viperKeys"
	utils_error "github.com/Uh-little-less-dum/go-utils/pkg/errorHandling"
	"github.com/go-git/go-git/v5"
	"github.com/spf13/viper"
)

type GitManager struct {
	Url       build_constants.BuildConstant
	Directory string
}

func (g GitManager) InitialClone(targetDir string) {
	err := os.Chdir(targetDir)
	utils_error.HandleError(err)
	timeout := time.Duration(viper.GetInt(string(viper_keys.CloneTimeout))) * time.Second
	ctx, cancelF := context.WithTimeout(context.Background(), timeout)
	defer cancelF()
	_, err = git.PlainCloneContext(ctx, g.Directory, false, &git.CloneOptions{
		URL:           string(build_constants.TemplateRepoUrl),
		RemoteName:    build_constants.TemplateRemoteName,
		ReferenceName: build_constants.TemplateBranch,
		SingleBranch:  true,
		// Progress:      os.Stdout,
	})
	utils_error.HandleErrorWithPrefix("Cloning: ", err)
	// VERSION_NEXT: This should eventually take in config arguments that will bind the cloned repo to a remote repo during the initial clone.
	// Remove the git directory to allow the user to setup their own git repo.
	// err = os.RemoveAll(filepath.Join(targetDir, ".git"))
	// utils_error.HandleError(err)
}

func NewTemplateAppGitManager(targetDirectory string) GitManager {
	return GitManager{
		Url:       build_constants.TemplateRepoUrl,
		Directory: targetDirectory,
	}
}
