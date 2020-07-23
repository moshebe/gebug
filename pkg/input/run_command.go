package input

import (
	"github.com/manifoldco/promptui"
	"github.com/moshebe/gebug/pkg/config"
)

// PromptRunCommand handles the prompt that asks for run command
type PromptRunCommand struct {
	*config.Config
}

// Run asks the user for the run command
func (p *PromptRunCommand) Run() error {
	prompt := &promptui.Prompt{
		Label:    "Run Command",
		Validate: nonEmptyValidator{field: &p.RunCommand}.validate,
		Default:  p.RunCommand,
	}
	_, err := prompt.Run()
	if err != nil {
		return err
	}
	return nil
}
