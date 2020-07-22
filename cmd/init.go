package cmd

import (
	"github.com/manifoldco/promptui"
	"github.com/moshebe/gebug/pkg/input"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a Gebug project",
	Long:  "Setup the relevant configuration files in order to run Gebug with this project",
	Run: func(cmd *cobra.Command, args []string) {
		currentConfig, _ := input.LoadOrDefault(workDir)

		err := input.Setup(currentConfig, []input.ConfigPrompt{
			&input.PromptName{Config: currentConfig},
			&input.PromptOutputBinary{Config: currentConfig},
			&input.PromptBuildCommand{Config: currentConfig},
			&input.PromptRunCommand{Config: currentConfig},
			&input.PromptRuntimeImage{Config: currentConfig},
			&input.PromptDebuggerOptions{Config: currentConfig},
			&input.PromptExposePort{Config: currentConfig},
			&input.PromptNetworks{Config: currentConfig},
		}, workDir)

		if err != nil {
			if err == promptui.ErrInterrupt {
				zap.L().Info("init command was interrupted by the user, ignoring unsaved changes...")
			} else {
				zap.L().Fatal("Failed to initialize project", zap.Error(err))
			}
		}
	},
}
