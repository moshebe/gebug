package input

import (
	"github.com/manifoldco/promptui"
	"github.com/moshebe/gebug/pkg/config"
)

type PromptEnvironment struct {
	*config.Config
}

func (p *PromptEnvironment) Run() error {
	prompt := NewSelectWithAddAndRemove(&p.Environment, &promptui.SelectWithAdd{
		Label:    "Define environment variables. Press existing choices to delete",
		AddLabel: "Add environment variable (e.g: FOO[=BAR])",
	})
	return prompt.Run()
}
