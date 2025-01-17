package cleanup_before_transpile_stages

import (
	"os"
	"os/exec"
	"path/filepath"
	"sync"

	build_config "github.com/Uh-little-less-dum/build/pkg/buildManager"
	viper_keys "github.com/Uh-little-less-dum/go-utils/pkg/constants/viperKeys"
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

func removeGitIfCleanGit(cfg *build_config.BuildManager) {
	if !viper.GetBool(string(viper_keys.CleanGit)) {
		return
	}
	err := os.RemoveAll(filepath.Join(cfg.TargetDir(), ".git"))
	if err != nil {
		log.Fatal(err)
	}
	c := exec.Command("git", "remote", "remove", "template")
	c.Dir = cfg.TargetDir()
	err = c.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func CleanupGit(cfg *build_config.BuildManager) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		removeGitIfCleanGit(cfg)
		defer wg.Done()
	}()
	go func() {
		WriteGitIgnore(cfg)
		defer wg.Done()
	}()
	wg.Wait()
}
