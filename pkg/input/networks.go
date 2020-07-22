package input

import (
	"github.com/manifoldco/promptui"
	"github.com/moshebe/gebug/pkg/config"
)

type PromptNetworks struct {
	*config.Config
}

func (p *PromptNetworks) Run() error {
	prompt := NewSelectWithAddAndRemove(&p.Networks, &promptui.SelectWithAdd{
		Label:    "Add networks or leave empty to create a new network",
		AddLabel: "Add network",
	})
	return prompt.Run()
}
