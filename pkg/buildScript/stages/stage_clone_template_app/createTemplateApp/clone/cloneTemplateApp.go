package stage_clone_template_app

import (
	git_manager "github.com/Uh-little-less-dum/build/pkg/buildScript/gitManager"
)

func Run(targetDir string) {
	gm := git_manager.NewTemplateAppGitManager(targetDir)
	gm.InitialClone(targetDir)
}

// func GetInitialCloneExecCommand(targetDir string) tea.ExecCommand {
// 	return charm_exec_command.NewCharmExecCommand(func() error {
// 		return Run(targetDir)
// 	})
// }

// func RunCmd(targetDir string) tea.Cmd {
// 	c := GetInitialCloneExecCommand(targetDir)
// 	return tea.Exec(c, func(err error) tea.Msg {
// 		return signals.SendFinishInitialTemplateCloneMsg()
// 	})
// }
