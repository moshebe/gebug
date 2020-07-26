package input

import (
	"fmt"
	"github.com/moshebe/gebug/pkg/setup"
	"github.com/moshebe/gebug/pkg/validate"
	"github.com/pkg/errors"
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

func (p *PromptDebuggerOptions) handleIde(name string, ide setup.Ide) error {
	detected, err := ide.Detected()
	if err != nil {
		return errors.WithMessage(err, "detect IDE in working directory")
	}

	if !detected {
		return nil
	}

	installed, err := ide.GebugInstalled()
	if err != nil {
		return errors.WithMessage(err, "check if Gebug is configured in IDE")
	}

	if installed {
		fmt.Printf("✔  Gebug is already configured in %s debugger\n", name)
		return nil
	}

	confirmPrompt := promptui.Prompt{
		Label:     fmt.Sprintf("IDE detected! would you like to configure Gebug in '%s'?", name),
		IsConfirm: true,
	}

	_, err = confirmPrompt.Run()
	if err == nil {
		zap.L().Debug("Configuring IDE", zap.String("name", name), zap.String("workDir", p.workDir))
		err = ide.Enable()
		if err != nil {
			return errors.WithMessage(err, "enable Gebug debugger configurations")
		}
		fmt.Printf("✔  Gebug configured in %s debugger successfully!\n", name)
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

	for name, ide := range setup.SupportedIde(p.workDir, p.DebuggerPort) {
		err := p.handleIde(name, ide)
		if err != nil {
			zap.L().Error("Failed to handle IDE checks", zap.String("name", name), zap.String("workDir", p.workDir))
		}

	}

	return nil
}
