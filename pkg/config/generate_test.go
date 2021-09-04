package config

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestConfig_Generate(t *testing.T) {
	AppFs = afero.NewMemMapFs()
	assertion := assert.New(t)

	workDir := "."
	err := mockConfig.Generate(workDir)
	assertion.NoError(err)

	dirExists, err := afero.DirExists(AppFs, RootDir)
	assertion.NoError(err)
	if !dirExists {
		err = AppFs.Mkdir(RootDir, 0777)
		assertion.NoError(err)
	}

	exists, err := afero.Exists(AppFs, FilePath(workDir, DockerfileName))
	assertion.NoError(err)
	assertion.True(exists)

	exists, err = afero.Exists(AppFs, FilePath(workDir, DockerComposeFileName))
	assertion.NoError(err)
	assertion.True(exists)
}
