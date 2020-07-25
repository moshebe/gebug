package setup

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"io"
	"io/ioutil"
	"path"
	"regexp"
)

var AppFs = afero.NewOsFs()

const launchConfigTemplate = `{
	"name": "Gebug",
	"type": "go",
	"request": "attach",
	"mode": "remote",
	"remotePath": "/src",
	"port": 40000,
	"host": "127.0.0.1"
}`

const (
	vscodeDirName        = ".vscode"
	vscodeLaunchFileName = "launch.json"
	gebugLaunchName      = "Gebug"
)

type VsCode struct {
	baseIde
}

func (v VsCode) installedInLaunchConfig(in []byte) (bool, error) {
	var launchConfig = struct {
		Version        string
		Configurations []struct {
			Name string
		}
	}{}

	re := regexp.MustCompile("(?s)//.*?\n|/\\*.*?\\*/")
	noCommentsInput := re.ReplaceAll(in, nil)

	err := json.Unmarshal(noCommentsInput, &launchConfig)
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

// AddGebugLaunchConfig edits the current `launch.json` content and append the `Gebug` debugging configuration
func (s *VsCode) AddGebugLaunchConfig(launchConfig []byte) (io.Writer, error) {

	return nil, nil
}
