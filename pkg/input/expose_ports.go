package input

import (
	"github.com/manifoldco/promptui"
	"github.com/moshebe/gebug/pkg/config"
)

type PromptExposePort struct {
	*config.Config
}

func (p *PromptExposePort) Run() error {
	prompt := NewSelectWithAddAndRemove(&p.ExposePorts, &promptui.SelectWithAdd{
		Label:    "Define ports to expose. Press existing choices to delete (e.g: 8080[:8080])",
		AddLabel: "Add port",
	})
	return prompt.Run()
}
