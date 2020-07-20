package input

import (
	"gebug/pkg/config"
	"github.com/manifoldco/promptui"
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
