package input

import (
	"github.com/manifoldco/promptui"
	"github.com/moshebe/gebug/pkg/config"
)

// PromptEnvironment handles the prompt that asks for envrionment variables
type PromptEnvironment struct {
	*config.Config
}

// Run asks the user for envrionment variables
func (p *PromptEnvironment) Run() error {
	prompt := NewSelectWithAddAndRemove(&p.Environment, &promptui.SelectWithAdd{
		Label:    "Define environment variables. Press existing choices to delete",
		AddLabel: "Add environment variable (e.g: FOO[=BAR])",
	})
	return prompt.Run()
}
