package input

import (
	"github.com/spf13/afero"
	"io/ioutil"
	"os"
	"path"

	"github.com/moshebe/gebug/pkg/osutil"

	"github.com/moshebe/gebug/pkg/config"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// AppFs hold the file-system abstraction for this package
var AppFs = afero.NewOsFs()

// ConfigPrompt asks for fields for the configuration
type ConfigPrompt interface {
	// Run asks for configuration field and saves it in configuration
	Run() error
}

// LoadOrDefault loads gebug's configuration file from the disk. Loads a default configuration in case of failure
func LoadOrDefault(workDir string) (*config.Config, bool) {
	fallback := &config.Config{
		OutputBinaryPath: "/app",
		BuildCommand:     `go build -o {{.output_binary}}`,
		RunCommand:       `{{.output_binary}}`,
		RuntimeImage:     "golang:1.14",
	}

	configFilePath := config.FilePath(workDir, config.Path)
	content, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		if !os.IsNotExist(err) {
			zap.L().Error("Failed to read configuration file", zap.String("path", configFilePath), zap.Error(err))
		}
		return fallback, false
	}
	cfg, err := config.Load(content)
	if err != nil {
		zap.L().Error("Failed to load configuration content", zap.Error(err))
		return fallback, false
	}

	return cfg, true
}

func save(workDir string, currentConfig *config.Config) error {
	if !osutil.FileExists(config.FilePath(workDir, config.Path)) {
		if !osutil.FileExists(path.Join(workDir, config.RootDir)) {
			err := os.Mkdir(path.Join(workDir, config.RootDir), 0755)
			if err != nil {
				return errors.WithMessage(err, "create config directory")
			}
		}
	}

	configFile, err := os.Create(config.FilePath(workDir, config.Path))
	if err != nil {
		return errors.WithMessage(err, "create config file")
	}
	defer configFile.Close()

	err = currentConfig.Write(configFile)
	if err != nil {
		return errors.WithMessage(err, "write configurations to config file")
	}

	return nil
}

// Setup runs a list of prompts and saves the user's input as config
func Setup(currentConfig *config.Config, prompts []ConfigPrompt, workDir string) error {
	for _, p := range prompts {
		err := p.Run()
		if err != nil {
			return err
		}
	}

	err := save(workDir, currentConfig)
	if err != nil {
		return errors.WithMessage(err, "save configuration")
	}

	return nil
}
