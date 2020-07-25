package input

import (
	"github.com/manifoldco/promptui"
	"github.com/moshebe/gebug/pkg/config"
)

// PromptOutputBinary handles the prompt that asks for output binary
type PromptOutputBinary struct {
	*config.Config
}

// Run asks the user for the name of the output binary
func (p *PromptOutputBinary) Run() error {
	prompt := &promptui.Prompt{
		Label:    "Output Binary Path (inside the container, referenced by: {{.output_binary}})",
		Validate: nonEmptyValidator{}.validate,
		Default:  p.OutputBinaryPath,
	}

	var err error
	p.OutputBinaryPath, err = prompt.Run()
	if err != nil {
		return err
	}

	return nil
}
