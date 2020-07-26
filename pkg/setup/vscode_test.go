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

var mockVsCode = &VsCode{baseIde{
	WorkDir:      ".",
	DebuggerPort: 43210,
}}

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
				AppFs = afero.NewMemMapFs()
				assertion := assert.New(t)
				if test.create {
					err := AppFs.Mkdir(vscodeDirName, 0777)
					assertion.NoError(err)
				}
				got, err := mockVsCode.Detected()
				assertion.NoError(err)
				assertion.Equal(test.expected, got)
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

func TestVsCode_GebugInstalled(t *testing.T) {
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
			AppFs = afero.NewMemMapFs()
			assertion := assert.New(t)
			filePath := path.Join("testdata", test.name+".in")
			input, err := ioutil.ReadFile(filePath)
			assertion.NoError(err)

			err = AppFs.Mkdir(vscodeDirName, 0777)
			assertion.NoError(err)
			err = afero.WriteFile(AppFs, mockVsCode.launchConfigFilePath(), input, 0777)
			assertion.NoError(err)
			got, err := mockVsCode.GebugInstalled()
			if test.wantErr {
				assertion.Error(err)
			} else {
				assertion.NoError(err)
			}
			assertion.Equal(test.expected, got)
		})
	}
}

func testEnableHelper(t *testing.T, input, golden *bytes.Buffer, f func() error) {
	assertion := assert.New(t)
	launchFilePath := mockVsCode.launchConfigFilePath()
	err := afero.WriteFile(AppFs, launchFilePath, input.Bytes(), 0777)
	assertion.NoError(err)

	err = f()
	assertion.NoError(err)

	got, err := afero.ReadFile(AppFs, launchFilePath)
	assertion.NoError(err)

	assertion.JSONEq(golden.String(), string(got))
}

func TestVsCode_Enable(t *testing.T) {
	testutil.RunTestData(t, "vscode_enable", func(t *testing.T, input, golden *bytes.Buffer) {
		AppFs = afero.NewMemMapFs()
		testEnableHelper(t, input, golden, mockVsCode.Enable)
	})
}

func TestVsCode_Disable(t *testing.T) {
	testutil.RunTestData(t, "vscode_disable", func(t *testing.T, input, golden *bytes.Buffer) {
		AppFs = afero.NewMemMapFs()
		testEnableHelper(t, input, golden, mockVsCode.Disable)
	})
}
