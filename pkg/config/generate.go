package config

import (
	"io"
	"path"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

const (
	// RootDir contains the name of the directory where gebug internals files and configurations are stored
	RootDir = ".gebug"

	// Path contains the name of gebug's configuration file
	Path = "config.yaml"

	// DockerfileName contains the name of the dockerfile that is used for building and running the project
	DockerfileName = "Dockerfile"

	// DockerComposeFileName contains the name of the docker-compose configuration that is used to configure the contianer
	DockerComposeFileName = "docker-compose.yml"
)

// FilePath returns the path of a file inside gebug's directory
func FilePath(workDir string, fileName string) string {
	return path.Join(workDir, RootDir, fileName)
}

func createConfigFile(fileName string, workDir string, renderFunc func(io.Writer) error) error {
	filePath := FilePath(workDir, fileName)
	zap.L().Debug("Generating config file", zap.String("path", filePath))
	file, err := AppFs.Create(filePath)
	if err != nil {
		return errors.WithMessagef(err, "create file '%s'", fileName)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			zap.L().Error("Failed to close file", zap.String("path", filePath))
		}
	}()

	err = renderFunc(file)
	if err != nil {
		return errors.WithMessagef(err, "generate file content: '%s'", fileName)
	}

	return nil
}

// Generate generates the required files to run the docker
func (c *Config) Generate(workDir string) error {
	for fileName, renderFunc := range map[string]func(io.Writer) error{
		DockerComposeFileName: c.RenderDockerComposeFile,
		DockerfileName:        c.RenderDockerfile,
	} {
		if err := createConfigFile(fileName, workDir, renderFunc); err != nil {
			return err
		}
	}
	return nil
}
