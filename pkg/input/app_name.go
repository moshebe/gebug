package input

import (
	"github.com/manifoldco/promptui"
	"github.com/moshebe/gebug/pkg/config"
	"github.com/moshebe/gebug/pkg/validate"
)

// PromptName handles the prompt that asks for application name
type PromptName struct {
	*config.Config
}

// Run asks the user for application name
func (p *PromptName) Run() error {
	prompt := &promptui.Prompt{
		Label:    "Application Name",
		Validate: validate.RegexValidator{Pattern: `^[A-Za-z0-9]([A-Za-z0-9_-]*[A-Za-z0-9])?$`}.Validate,
		Default:  p.Name,
	}

	var err error
	p.Name, err = prompt.Run()
	if err != nil {
		return err
	}

	return nil
}
