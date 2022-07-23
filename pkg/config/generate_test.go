package config

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
)

func TestConfig_Generate(t *testing.T) {
	AppFs = afero.NewMemMapFs()

	workDir := "."
	err := mockConfig.Generate(workDir)
	require.NoError(t, err)

	dirExists, err := afero.DirExists(AppFs, RootDir)
	require.NoError(t, err)
	if !dirExists {
		err = AppFs.Mkdir(RootDir, 0777)
		require.NoError(t, err)
	}

	exists, err := afero.Exists(AppFs, FilePath(workDir, DockerfileName))
	require.NoError(t, err)
	require.True(t, exists)

	exists, err = afero.Exists(AppFs, FilePath(workDir, DockerComposeFileName))
	require.NoError(t, err)
	require.True(t, exists)
}
