package input

import (
	"fmt"
	"github.com/moshebe/gebug/pkg/setup"
	"github.com/moshebe/gebug/pkg/validate"
	"go.uber.org/zap"
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
	workDir string
}

func (p *PromptDebuggerOptions) autoSetupIde() error {
	supportedIde := map[string]setup.Ide{
		"Visual Studio Code": setup.NewVsCode(p.workDir, p.DebuggerPort),
	}

	for name, ide := range supportedIde {
		detected, err := ide.Detected()
		if err != nil {
			zap.L().Error("Failed to detect IDE in working directory", zap.String("workDir", p.workDir),
				zap.String("IDE", name), zap.Error(err))
			continue
		}

		if !detected {
			continue
		}

		// TODO: PROMPT CONFIRMATION OF ENABLE IDE DEBUGGER OPTIONS

	}

	return nil
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

	p.DebuggerEnabled = value == debugger
	if !p.DebuggerEnabled {
		return nil
	}

	prompt := &promptui.Prompt{
		Label: "Debugger Port",
		Validate: validate.NumericRangeValidator{
			Min: 1024,
			Max: 65535,
		}.Validate,
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

	// TODO: check for IDE

	return nil
}
