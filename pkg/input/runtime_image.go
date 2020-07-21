package input

import (
	"github.com/manifoldco/promptui"
	"github.com/moshebe/gebug/pkg/config"
)

type PromptRuntimeImage struct {
	*config.Config
}

func (p *PromptRuntimeImage) Run() error {
	prompt := &promptui.Prompt{
		Label:    "Runtime Docker Image",
		Validate: NonEmptyValidator{field: &p.RuntimeImage}.validate,
		Default:  p.RuntimeImage,
	}
	_, err := prompt.Run()
	if err != nil {
		return err
	}
	return nil
}
