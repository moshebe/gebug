package setup

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestIde_detected(t *testing.T) {
	AppFs = afero.NewMemMapFs()
	assertion := assert.New(t)

	createdDirName := ".my_ide"
	err := AppFs.Mkdir(createdDirName, 0777)
	assertion.NoError(err)

	_, err = afero.ReadDir(AppFs, createdDirName)
	assertion.NoError(err)

	_, err = afero.ReadDir(AppFs, ".not-exists")
	assertion.Error(err)
}

func TestIde_SupportedIdes(t *testing.T) {
	assert.NotEmpty(t, SupportedIdes(".", 0))
}
