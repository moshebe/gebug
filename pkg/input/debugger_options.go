package input

import (
	"fmt"
	"gebug/pkg/config"
	"github.com/manifoldco/promptui"
)

type PromptDebuggerOptions struct {
	*config.Config
}

func (p *PromptDebuggerOptions) Run() error {
	selectPrompt := promptui.Select{
		Label: "Select debugging method",
		Items: []string{"Hot Reload", "Debugger"},
	}
	index, _, err := selectPrompt.Run()
	if err != nil {
		return err
	}

	p.DebuggerEnabled = index == 1
	if p.DebuggerEnabled {
		prompt := &promptui.Prompt{
			Label: "Debugger Port",
			Validate: NumericRangeValidator{
				min:   1024,
				max:   65535,
				field: &p.DebuggerPort,
			}.validate,
			Default: fmt.Sprintf("%d", p.DebuggerPort),
		}
		_, err := prompt.Run()
		if err != nil {
			return err
		}
	}
	return nil
}
