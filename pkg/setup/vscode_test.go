package setup

import (
	"bytes"
	"github.com/moshebe/gebug/pkg/testutil"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"path"
	"strconv"
	"testing"
)

var mockVsCode = &VsCode{baseIde{workDir: "."}}

func TestVscodeStuff(t *testing.T) {
	assertion := assert.New(t)

	// TODO: maybe use Setup function to clear FS
	AppFs = afero.NewMemMapFs()

	err := AppFs.Mkdir(".vscode", 0777)
	assertion.NoError(err)

	results, err := afero.ReadDir(AppFs, ".vscode2")
	assertion.NoError(err)

	t.Log(results)
}

func TestVsCode_Detected(t *testing.T) {
	tests := []struct {
		create   bool
		expected bool
	}{
		{create: true, expected: true},
		{create: false, expected: false},
	}
	{
		for i, test := range tests {
			t.Run(strconv.Itoa(i), func(t *testing.T) {
				testutil.FsTest(t, &AppFs, func(t *testing.T) {
					assertion := assert.New(t)
					if test.create {
						err := AppFs.Mkdir(vscodeDirName, 0777)
						assertion.NoError(err)
					}
					got, err := mockVsCode.Detected()
					assertion.NoError(err)
					assertion.Equal(test.expected, got)
				})
			})
		}
	}
}

func TestVsCode_installedInLaunchConfig(t *testing.T) {
	tests := []struct {
		name     string
		expected bool
		wantErr  bool
	}{
		{
			name:     "vscode_install_in_launch_config_0_not_exists",
			expected: false,
			wantErr:  false,
		},
		{
			name:     "vscode_install_in_launch_config_1_exists",
			expected: true,
			wantErr:  false,
		},
		{
			name:     "vscode_install_in_launch_config_2_exists_both",
			expected: true,
			wantErr:  false,
		},
		{
			name:     "vscode_install_in_launch_config_3_empty",
			expected: false,
			wantErr:  true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assertion := assert.New(t)
			filePath := path.Join("testdata", test.name+".in")
			input, err := ioutil.ReadFile(filePath)
			assertion.NoError(err)

			got, err := mockVsCode.installedInLaunchConfig(bytes.NewBuffer(input).Bytes())
			if test.wantErr {
				assertion.Error(err)
			} else {
				assertion.NoError(err)
			}
			assertion.Equal(test.expected, got)
		})
	}

}
