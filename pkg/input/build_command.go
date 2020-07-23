package input

import (
	"github.com/manifoldco/promptui"
	"github.com/moshebe/gebug/pkg/config"
)

// PromptBuildCommand handles the prompt that asks for build command
type PromptBuildCommand struct {
	*config.Config
}

// Run asks the user for build command
func (p *PromptBuildCommand) Run() error {
	prompt := &promptui.Prompt{
		Label:    "Build Command",
		Validate: nonEmptyValidator{field: &p.BuildCommand}.validate,
		Default:  p.BuildCommand,
	}
	_, err := prompt.Run()
	if err != nil {
		return err
	}
	return nil
}
