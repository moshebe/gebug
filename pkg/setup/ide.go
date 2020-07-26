package setup

import (
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"path"
)

var AppFs = afero.NewOsFs()

type Ide interface {
	// Detected tells if the IDE trails were found in the working directory. e.g: `.vscode` or `.idea` directories.
	Detected() (bool, error)

	// GebugInstalled tells if Gebug debugger mode was set in the IDE
	GebugInstalled() (bool, error)

	// Enable Gebug's debugger configurations
	Enable() error

	// Disable Gebug's debugger configurations
	Disable() error
}

type baseIde struct {
	WorkDir      string
	DebuggerPort int
}

func (i baseIde) detected(ideDirName string) (bool, error) {
	detected, err := afero.DirExists(AppFs, path.Join(i.WorkDir, ideDirName))
	if err != nil {
		return false, errors.WithMessage(err, "check if directory exists")
	}

	return detected, nil
}

func SupportedIdes(workDir string, port int) map[string]Ide {
	return map[string]Ide{
		"Visual Studio Code": &VsCode{baseIde{WorkDir: workDir, DebuggerPort: port}},
	}
}
