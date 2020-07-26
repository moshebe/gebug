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

var mockVsCode = &VsCode{baseIde{workDir: ".", debuggerPort: 43210,}}

// TODO: remove
//func TestVscodeStuff(t *testing.T) {
//	assertion := assert.New(t)
//
//	input, err := ioutil.ReadFile("testdata/vscode_test_file.in")
//	assertion.NoError(err)
//	t.Log(string(input))
//	// TODO: maybe use Setup function to clear FS
//	AppFs = afero.NewMemMapFs()
//
//	//err = AppFs.Mkdir(".vscode", 0777)
//	//assertion.NoError(err)
//
//	//	content := `{
//	//    "version": "0.2.0",
//	//    "configurations": [
//	//        {
//	//            "name": "Launch",
//	//            "type": "go",
//	//            "request": "launch",
//	//            "mode": "auto",
//	//            "program": "${fileDirname}",
//	//            "env": {},
//	//            "args": []
//	//        }
//	//    ]
//	//}`
//
//	//err = afero.WriteFile(AppFs, path.Join(".vscode", "launch.json"),
//	//	[]byte(content), 0777)
//	//assertion.NoError(err)
//	//var msg json.RawMessage
//
//	var launchConfig = struct {
//		Version        string
//		Configurations []interface{}
//	}{}
//
//	//re := regexp.MustCompile("(?s)//.*?\n|/\\*.*?\\*/")
//	//noCommentsInput := re.ReplaceAll(in, nil)
//	//
//	err = json.Unmarshal(input, &launchConfig)
//	assertion.NoError(err)
//	launchConfig.Configurations = append(launchConfig.Configurations, mockVsCode.createGebugConfig())
//
//
//	output, err := mockVsCode.editLaunchConfig(true, input)
//	assertion.NoError(err)
//
//	t.Log(string(output))
//	//if err != nil {
//	//	return false, errors.WithMessage(err, "unmarshal no comments input")
//	//}
//
//}

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
		testutil.FsTest(t, &AppFs, func(t *testing.T) {
			testEnableHelper(t, input, golden, mockVsCode.Enable)
		})
	})
}

func TestVsCode_Disable(t *testing.T) {
	testutil.RunTestData(t, "vscode_disable", func(t *testing.T, input, golden *bytes.Buffer) {
		testutil.FsTest(t, &AppFs, func(t *testing.T) {
			testEnableHelper(t, input, golden, mockVsCode.Disable)
		})
	})
}
