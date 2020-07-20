package input

import (
	"gebug/pkg/config"
	"github.com/manifoldco/promptui"
)

type PromptRunCommand struct {
	*config.Config
}

func (p *PromptRunCommand) Run() error {
	prompt := &promptui.Prompt{
		Label:    "Run Command",
		Validate: NonEmptyValidator{field: &p.RunCommand}.validate,
		Default:  p.RunCommand,
	}
	_, err := prompt.Run()
	if err != nil {
		return err
	}
	return nil
}
