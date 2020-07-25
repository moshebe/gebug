package input

import (
	"github.com/manifoldco/promptui"
	"github.com/moshebe/gebug/pkg/config"
)

// PromptRuntimeImage handles the prompt that asks for runtime image
type PromptRuntimeImage struct {
	*config.Config
}

// Run asks the user for the runtime image
func (p *PromptRuntimeImage) Run() error {
	prompt := &promptui.Prompt{
		Label:    "Runtime Docker Image",
		Validate: nonEmptyValidator{}.validate,
		Default:  p.RuntimeImage,
	}

	var err error
	p.RuntimeImage, err = prompt.Run()
	if err != nil {
		return err
	}

	return nil
}
