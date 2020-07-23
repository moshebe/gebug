package input

import (
	"github.com/manifoldco/promptui"
	"github.com/moshebe/gebug/pkg/config"
)

// PromptExposePort handles the prompt that asks for expose ports
type PromptExposePort struct {
	*config.Config
}

// Run asks the user for expose ports
func (p *PromptExposePort) Run() error {
	prompt := NewSelectWithAddAndRemove(&p.ExposePorts, &promptui.SelectWithAdd{
		Label:    "Define ports to expose. Press existing choices to delete (e.g: 8080[:8080])",
		AddLabel: "Add port",
	})
	return prompt.Run()
}
