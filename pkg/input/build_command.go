package input

import (
	"github.com/manifoldco/promptui"
	"github.com/moshebe/gebug/pkg/config"
)

type PromptBuildCommand struct {
	*config.Config
}

func (p *PromptBuildCommand) Run() error {
	prompt := &promptui.Prompt{
		Label:    "Build Command",
		Validate: NonEmptyValidator{field: &p.BuildCommand}.validate,
		Default:  p.BuildCommand,
	}
	_, err := prompt.Run()
	if err != nil {
		return err
	}
	return nil
}
