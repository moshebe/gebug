package setup

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
)

func TestIde_detected(t *testing.T) {
	AppFs = afero.NewMemMapFs()

	createdDirName := ".my_ide"
	err := AppFs.Mkdir(createdDirName, 0777)
	require.NoError(t, err)

	_, err = afero.ReadDir(AppFs, createdDirName)
	require.NoError(t, err)

	_, err = afero.ReadDir(AppFs, ".not-exists")
	require.Error(t, err)
}

func TestIde_SupportedIdes(t *testing.T) {
	require.NotEmpty(t, SupportedIdes(".", 0))
}
