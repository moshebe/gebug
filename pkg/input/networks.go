package input

import (
	"github.com/manifoldco/promptui"
	"github.com/moshebe/gebug/pkg/config"
)

// PromptNetworks handles the prompt that asks for networks
type PromptNetworks struct {
	*config.Config
}

// Run asks the user for networks to use
func (p *PromptNetworks) Run() error {
	prompt := NewSelectWithAddAndRemove(&p.Networks, &promptui.SelectWithAdd{
		Label:    "Add networks or leave empty to create a new network",
		AddLabel: "Add network",
	})
	return prompt.Run()
}
