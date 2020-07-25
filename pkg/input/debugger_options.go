package input

import (
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/moshebe/gebug/pkg/config"
)

const (
	hotReload = "Hot Reload"
	debugger  = "Debugger"
)

// PromptDebuggerOptions handles the prompt that asks for debugger options
type PromptDebuggerOptions struct {
	*config.Config
}

// Run asks the user for debugger options
func (p *PromptDebuggerOptions) Run() error {
	selectPrompt := promptui.Select{
		Label: "Select debugging method",
		Items: []string{hotReload, debugger},
	}
	_, value, err := selectPrompt.Run()
	if err != nil {
		return err
	}

	p.DebuggerEnabled = (value == debugger)
	if !p.DebuggerEnabled {
		return nil
	}

	prompt := &promptui.Prompt{
		Label: "Debugger Port",
		Validate: numericRangeValidator{
			min: 1024,
			max: 65535,
		}.validate,
		Default: fmt.Sprintf("%d", p.DebuggerPort),
	}

	value, err = prompt.Run()
	if err != nil {
		return err
	}

	p.DebuggerPort, err = strconv.Atoi(value)
	if err != nil {
		return err
	}

	return nil
}
