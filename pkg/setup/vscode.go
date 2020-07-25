package setup

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"io/ioutil"
	"os"
	"path"
	"regexp"
)

const (
	vscodeDirName        = ".vscode"
	vscodeLaunchFileName = "launch.json"
	gebugLaunchName      = "Gebug"
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
	launchConfigPath := path.Join(v.workDir, vscodeDirName, vscodeLaunchFileName)
	launchContent, err := ioutil.ReadFile(launchConfigPath)
	if err != nil {
		return false, errors.WithMessage(err, "read vscode launch.json file")
	}
	installed, err := v.installedInLaunchConfig(launchContent)
	if err != nil {
		return false, errors.WithMessage(err, "check if installed in launch.json")
	}

	return installed, nil
}

// TODO: tests -
// if no file - create one and put gebug inside
// has file, gebug not found - add it
// has file, has gebug - don't change anything
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

func (v VsCode) removeComments(in [] byte) []byte {
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
		"port":       v.debuggerPort,
	}
}

func (v VsCode) editConfig(enabled bool, input []byte) ([]byte, error) {
	var launchConfig = struct {
		Version        string
		Configurations []interface{}
	}{}
	err := json.Unmarshal(input, &launchConfig)
	if err != nil {
		return nil, errors.WithMessage(err, "unmarshal launch.json")
	}

	var newConfig []interface{}
	if enabled {
		newConfig = append(launchConfig.Configurations, v.createGebugConfig())
	} else {
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
	}
	launchConfig.Configurations = newConfig

	output, err := json.MarshalIndent(launchConfig, "", "\t")
	if err != nil {
		return nil, errors.WithMessage(err, "marshal new configuration")
	}

	return output, nil
}

func (v VsCode) launchConfigFilePath() string {
	return path.Join(v.workDir, vscodeDirName, vscodeLaunchFileName)
}

func (v VsCode) setEnabled(enabled bool) error {
	launchFilePath := v.launchConfigFilePath()

	perm := os.FileMode(0755)
	file, err := AppFs.Stat(launchFilePath)
	if err != nil {
		if !os.IsNotExist(err) {
			return errors.WithMessage(err, "access launch.json config file")
		} else {
			// no file is ok in case of disable
			if !enabled {
				return nil
			}
		}
	} else {
		perm = file.Mode()
	}

	input, err := afero.ReadFile(AppFs, launchFilePath)
	if err != nil {
		return errors.WithMessage(err, "read launch.json file")
	}

	output, err := v.editConfig(true, input)
	if err != nil {
		return errors.WithMessage(err, "edit configuration to set enabled mode")
	}

	err = afero.WriteFile(AppFs, launchFilePath, output, perm)
	if err != nil {
		return errors.WithMessage(err, "write configuration to launch.json")
	}

	return nil
}
