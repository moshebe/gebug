package input

import (
	"github.com/manifoldco/promptui"
	"github.com/moshebe/gebug/pkg/config"
	"github.com/moshebe/gebug/pkg/validate"
)

// PromptRunCommand handles the prompt that asks for run command
type PromptRunCommand struct {
	*config.Config
}

// Run asks the user for the run command
func (p *PromptRunCommand) Run() error {
	prompt := &promptui.Prompt{
		Label:    "Run Command",
		Validate: validate.NonEmptyValidator{}.Validate,
		Default:  p.RunCommand,
	}

	var err error
	p.RunCommand, err = prompt.Run()
	if err != nil {
		return err
	}
	return nil
}
