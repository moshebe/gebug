package input

import (
	"github.com/manifoldco/promptui"
	"github.com/moshebe/gebug/pkg/config"
)

type PromptOutputBinary struct {
	*config.Config
}

func (p *PromptOutputBinary) Run() error {
	prompt := &promptui.Prompt{
		Label:    "Output Binary Path (inside the container, referenced by: {{.output_binary}})",
		Validate: NonEmptyValidator{field: &p.OutputBinaryPath}.validate,
		Default:  p.OutputBinaryPath,
	}

	_, err := prompt.Run()
	if err != nil {
		return err
	}

	return nil
}
