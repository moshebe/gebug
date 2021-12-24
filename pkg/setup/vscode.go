package setup

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"regexp"

	"github.com/spf13/afero"
)

const (
	vscodeDirName            = ".vscode"
	vscodeLaunchFileName     = "launch.json"
	gebugLaunchName          = "Gebug"
	defaultVsCodeConfVersion = "0.2.0"
)

// VsCode is the 'Visual Studio Code' integration with Gebug
type VsCode struct {
	baseIde
}

// Detected tells if the IDE trails were found in the working directory (the `.vscode` directory exists)
func (v VsCode) Detected() (bool, error) {
	return v.detected(vscodeDirName)
}

// GebugInstalled tells if Gebug debugger mode was set in the IDE 'launch.json' file
func (v VsCode) GebugInstalled() (bool, error) {
	detected, err := v.Detected()
	if err != nil {
		return false, err
	}
	if !detected {
		return false, fmt.Errorf("vscode was not detected in this workspace")
	}
	launchConfigPath := path.Join(v.WorkDir, vscodeDirName, vscodeLaunchFileName)
	launchContent, err := afero.ReadFile(AppFs, launchConfigPath)
	if err != nil {
		return false, fmt.Errorf("read vscode launch.json file: %w", err)
	}
	installed, err := v.installedInLaunchConfig(launchContent)
	if err != nil {
		return false, fmt.Errorf("check if installed in launch.json: %w", err)
	}

	return installed, nil
}

// Enable Gebug's debugger configurations by adding the Gebug object from the configurations json in 'launch.json'
func (v VsCode) Enable() error {
	return v.setEnabled(true)
}

// Disable Gebug's debugger configurations by removing the Gebug object from the configurations json in 'launch.json'
func (v VsCode) Disable() error {
	return v.setEnabled(false)
}

func (v VsCode) installedInLaunchConfig(in []byte) (bool, error) {
	var launchConfig = struct {
		Version        string
		Configurations []struct {
			Name string
		}
	}{}

	err := json.Unmarshal(v.removeComments(in), &launchConfig)
	if err != nil {
		return false, fmt.Errorf("unmarshal no comments input: %w", err)
	}
	for _, c := range launchConfig.Configurations {
		if c.Name != gebugLaunchName {
			continue
		}
		return true, nil
	}
	return false, nil
}

func (v VsCode) removeComments(in []byte) []byte {
	re := regexp.MustCompile("(?s)//.*?\n|/\\*.*?\\*/")
	return re.ReplaceAll(in, nil)
}

func (v VsCode) createGebugConfig() map[string]interface{} {
	return map[string]interface{}{
		"name":       gebugLaunchName,
		"type":       "go",
		"request":    "attach",
		"mode":       "remote",
		"remotePath": "/src",
		"host":       "127.0.0.1",
		"port":       v.DebuggerPort,
	}
}

func (v VsCode) editLaunchConfig(enabled bool, input []byte) ([]byte, error) {
	var launchConfig = struct {
		Version        string        `json:"version"`
		Configurations []interface{} `json:"configurations"`
	}{}
	// first, we remove the Gebug configuration if exists and then we add the new one if enabled
	newConfig := make([]interface{}, 0)

	if len(input) > 0 {
		err := json.Unmarshal(v.removeComments(input), &launchConfig)
		if err != nil {
			return nil, fmt.Errorf("unmarshal launch.json: %w", err)
		}

		for _, c := range launchConfig.Configurations {
			item, ok := c.(map[string]interface{})
			if !ok {
				return nil, fmt.Errorf("parse launch configuration")
			}
			if item["name"] == gebugLaunchName {
				continue
			}
			newConfig = append(newConfig, c)
		}
	} else {
		launchConfig.Version = defaultVsCodeConfVersion
	}

	if enabled {
		newConfig = append(newConfig, v.createGebugConfig())
	}
	launchConfig.Configurations = newConfig
	output, err := json.MarshalIndent(launchConfig, "", "\t")
	if err != nil {
		return nil, fmt.Errorf("marshal new configuration: %w", err)
	}

	return output, nil
}

func (v VsCode) launchConfigFilePath() string {
	return path.Join(v.WorkDir, vscodeDirName, vscodeLaunchFileName)
}

func (v VsCode) setEnabled(enabled bool) error {
	launchFilePath := v.launchConfigFilePath()

	perm := os.FileMode(0755)
	file, err := AppFs.Stat(launchFilePath)
	if err != nil {
		if !os.IsNotExist(err) {
			return fmt.Errorf("access launch.json config file: %w", err)
		}

		// no file is ok in case of disable
		if !enabled {
			return nil
		}

	} else {
		perm = file.Mode()
	}

	input, err := afero.ReadFile(AppFs, launchFilePath)
	if err != nil {
		return fmt.Errorf("read launch.json file: %w", err)
	}

	output, err := v.editLaunchConfig(enabled, input)
	if err != nil {
		return fmt.Errorf("edit configuration to set enabled mode: %w", err)
	}

	err = afero.WriteFile(AppFs, launchFilePath, output, perm)
	if err != nil {
		return fmt.Errorf("write configuration to launch.json: %w", err)
	}

	return nil
}
