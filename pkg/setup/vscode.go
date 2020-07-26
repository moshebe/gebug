package setup

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"os"
	"path"
	"regexp"
)

const (
	vscodeDirName            = ".vscode"
	vscodeLaunchFileName     = "launch.json"
	gebugLaunchName          = "Gebug"
	defaultVsCodeConfVersion = "0.2.0"
)

type VsCode struct {
	baseIde
}

func (v VsCode) Detected() (bool, error) {
	return v.detected(vscodeDirName)
}

func (v VsCode) GebugInstalled() (bool, error) {
	detected, err := v.Detected()
	if err != nil {
		return false, err
	}
	if !detected {
		return false, errors.New("vscode was not detected in this workspace")
	}
	launchConfigPath := path.Join(v.WorkDir, vscodeDirName, vscodeLaunchFileName)
	launchContent, err := afero.ReadFile(AppFs, launchConfigPath)
	if err != nil {
		return false, errors.WithMessage(err, "read vscode launch.json file")
	}
	installed, err := v.installedInLaunchConfig(launchContent)
	if err != nil {
		return false, errors.WithMessage(err, "check if installed in launch.json")
	}

	return installed, nil
}

func (v VsCode) Enable() error {
	return v.setEnabled(true)
}

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
		return false, errors.WithMessage(err, "unmarshal no comments input")
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
			return nil, errors.WithMessage(err, "unmarshal launch.json")
		}

		for _, c := range launchConfig.Configurations {
			item, ok := c.(map[string]interface{})
			if !ok {
				return nil, errors.New("parse launch configuration")
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
		return nil, errors.WithMessage(err, "marshal new configuration")
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
			return errors.WithMessage(err, "access launch.json config file")
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
		return errors.WithMessage(err, "read launch.json file")
	}

	output, err := v.editLaunchConfig(enabled, input)
	if err != nil {
		return errors.WithMessage(err, "edit configuration to set enabled mode")
	}

	err = afero.WriteFile(AppFs, launchFilePath, output, perm)
	if err != nil {
		return errors.WithMessage(err, "write configuration to launch.json")
	}

	return nil
}
