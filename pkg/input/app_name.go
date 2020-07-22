package input

import (
	"github.com/manifoldco/promptui"
	"github.com/moshebe/gebug/pkg/config"
)

// PromptName handles the prompt that asks for application name
type PromptName struct {
	*config.Config
}

// Run asks the user for input and saves the result
func (p *PromptName) Run() error {
	prompt := &promptui.Prompt{
		Label:    "Application Name",
		Validate: RegexValidator{`^[A-Za-z0-9]([A-Za-z0-9_-]*[A-Za-z0-9])?$`, &p.Name}.validate,
		Default:  p.Name,
	}
	_, err := prompt.Run()
	if err != nil {
		return err
	}
	return nil
}
