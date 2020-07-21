package input

import (
	"github.com/manifoldco/promptui"
	"github.com/moshebe/gebug/pkg/config"
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
