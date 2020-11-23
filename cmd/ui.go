package cmd

import (
	"fmt"
	"github.com/moshebe/gebug/pkg/osutil"
	"github.com/moshebe/gebug/pkg/web"
	"github.com/pkg/browser"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"sync"
)

var projectPath string
var uiPort int

const (
	imageName     = "gebug-ui"
	defaultUiPort = 3030
)

func init() {
	uiCmd.PersistentFlags().StringVar(&projectPath, "path", ".", "project directory path")
	uiCmd.PersistentFlags().IntVar(&uiPort, "port", defaultUiPort, "web UI port")

	rootCmd.AddCommand(uiCmd)
}

var uiCmd = &cobra.Command{
	Use:   "ui",
	Short: "Start Gebug web UI",
	Long:  "Start web UI that manages the Gebug configurations",
	Run: func(cmd *cobra.Command, args []string) {
		zap.L().Info("🚀 Launching Gebug web UI...")

		if projectPath == "." || projectPath == "" {
			zap.L().Debug("Resolving current project path", zap.String("projectPath", projectPath))
			cwd, err := os.Getwd()
			if err != nil {
				zap.L().Fatal("Failed to get current working directory", zap.Error(err))
			}
			projectPath = cwd
		}

		file, err := ioutil.TempFile("", "gebug-webui-docker-compose.*.yml")
		if err != nil {
			zap.L().Fatal("Failed to create temporary file for generating docker-compose", zap.Error(err))
		}

		zap.L().Debug("Generating docker-compose configuration", zap.String("path", file.Name()))

		err = web.RenderDockerCompose(&web.Opts{
			ImageName: imageName,
			Port:      uiPort,
			Location:  projectPath,
		}, file)

		if err != nil {
			zap.L().Fatal("Failed to generate configuration for docker-compose", zap.Error(err))
		}

		prerequisites := []string{"docker", "docker-compose"}
		for _, bin := range prerequisites {
			if !osutil.CommandExists(bin) {
				zap.S().Fatalf("'%s' was not found, please make sure it is installed correctly...", bin)
			}
		}

		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()
			if err := osutil.RunCommand(fmt.Sprintf("docker-compose -f %s up", file.Name())); err != nil {
				zap.L().Fatal("Failed to start web ui", zap.Error(err))
			}
		}()

		url := fmt.Sprintf("http://localhost:%d", uiPort)
		err = web.ReadinessProbe(url, verbose)
		if err != nil {
			zap.L().Fatal("Failed to start web ui server", zap.String("url", url), zap.Error(err))
		}
		zap.S().Infof("🔗 Ready on %s", url)
		err = browser.OpenURL(url)
		if err != nil {
			zap.L().Error("Failed to launch browser with url", zap.String("url", url), zap.Error(err))
		}
		wg.Wait()
		zap.L().Info("Finished")
	},
}
