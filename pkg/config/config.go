package config

import (
	"github.com/spf13/afero"
	"io"
	"path"
	"reflect"
	"strings"

	"github.com/moshebe/gebug/pkg/render"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// AppFs hold the file-system abstraction for this package
var AppFs = afero.NewOsFs()

// Config contains the fields of gebug configuration
type Config struct {
	Name             string   `yaml:"name" json:"name"`
	OutputBinaryPath string   `yaml:"output_binary" json:"output_binary"`
	BuildCommand     string   `yaml:"build_command" json:"build_command"`
	RunCommand       string   `yaml:"run_command" json:"run_command"`
	RuntimeImage     string   `yaml:"runtime_image" json:"runtime_image"`
	DebuggerEnabled  bool     `yaml:"debugger_enabled" json:"debugger_enabled"`
	DebuggerPort     int      `yaml:"debugger_port" json:"debugger_port"`
	ExposePorts      []string `yaml:"expose_ports" json:"expose_ports"`
	Networks         []string `yaml:"networks" json:"networks"`
	Environment      []string `yaml:"environment" json:"environment"`
}

func updateBuildCommand(buildCommand string, debuggerEnabled bool) string {
	buildCommand = strings.TrimSpace(buildCommand)
	if !debuggerEnabled || strings.Contains(buildCommand, "gcflags") {
		return buildCommand
	}

	goBuildPrefix := "go build"
	commandArgs := strings.TrimPrefix(buildCommand, goBuildPrefix)
	return strings.TrimSpace(goBuildPrefix + ` -gcflags="all=-N -l"` + commandArgs)
}

// Load loads a configuration to a Config struct
func Load(input []byte) (*Config, error) {
	c := &Config{}
	err := yaml.Unmarshal(input, c)
	if err != nil {
		return nil, errors.WithMessage(err, "unmarshal configuration")
	}

	err = c.render()
	if err != nil {
		return nil, errors.WithMessage(err, "render configuration")
	}

	c.BuildCommand = updateBuildCommand(c.BuildCommand, c.DebuggerEnabled)
	return c, nil
}

func ResolvePath(projectPath string) string {
	if strings.HasSuffix(projectPath, Path) {
		return projectPath
	}

	if strings.HasSuffix(projectPath, RootDir) || strings.HasSuffix(projectPath, RootDir+"/") {
		return path.Join(projectPath, Path)
	}

	return FilePath(projectPath, Path)
}

func (c Config) Write(writer io.Writer) error {
	out, err := yaml.Marshal(c)
	if err != nil {
		return errors.WithMessage(err, "marshal configuration")
	}
	_, err = writer.Write(out)
	if err != nil {
		return errors.WithMessage(err, "write marshalled configuration")
	}
	return nil
}

func (c *Config) render() error {
	values := map[string]string{}

	typeElement := reflect.TypeOf(c).Elem()
	valueElement := reflect.ValueOf(c).Elem()
	fields := typeElement.NumField()

	for i := 0; i < fields; i++ {
		tagName := typeElement.Field(i).Tag.Get("yaml")
		values[tagName] = valueElement.Field(i).String()
	}

	for i := 0; i < fields; i++ {
		fieldType := typeElement.Field(i)
		if fieldType.Type.String() != "string" {
			continue
		}

		rawValue := valueElement.Field(i).String()
		newValue, err := render.Render(rawValue, values)
		if err != nil {
			return errors.Errorf("unable to render value: '%s'", rawValue)
		}
		valueElement.Field(i).SetString(newValue)
	}
	return nil
}
