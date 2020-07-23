package input

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/moshebe/gebug/pkg/config"
)

// PromptDebuggerOptions handles the prompt that asks for debugger options
type PromptDebuggerOptions struct {
	*config.Config
}

// Run asks the user for debugger options
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
			Validate: numericRangeValidator{
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
