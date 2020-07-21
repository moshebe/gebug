package cmd

import (
	"fmt"
	"path"

	"github.com/moshebe/gebug/pkg/config"
	"github.com/moshebe/gebug/pkg/input"
	"github.com/moshebe/gebug/pkg/osutil"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var skipGenerate bool

func init() {
	runCmd.PersistentFlags().BoolVar(&skipGenerate, "skip-generate", false, "skip configuration file generation before start")
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "start",
	Short: "Start Gebug services",
	Long:  "Start the Docker containers which handles the debugging methods for the application",
	Run: func(cmd *cobra.Command, args []string) {
		currentConfig, ok := input.LoadOrDefault(workDir)
		if !ok {
			zap.L().Fatal("Failed to load configuration, are you sure you setup everything correctly? I suggest running `gebug init`...")
		}
		if !skipGenerate {
			zap.L().Debug("Skipping generation of configuration file")
			err := currentConfig.Generate(workDir)
			if err != nil {
				zap.L().Fatal("Failed to generate configuration files for current config", zap.Error(err))
			}
		}
		requiredFiles := []string{
			path.Join(workDir, config.RootDir),
			config.FilePath(workDir, config.Path),
			config.FilePath(workDir, config.DockerfileName),
			config.FilePath(workDir, config.DockerComposeFileName),
		}
		for _, file := range requiredFiles {
			if !osutil.FileExists(file) {
				zap.S().Fatalf("missing required file: '%s', please make sure you have generated it...", file)
			}
		}
		prerequisites := []string{"docker", "docker-compose"}
		for _, bin := range prerequisites {
			if !osutil.CommandExists(bin) {
				zap.S().Fatalf("'%s' was not found, please make sure it is installed correctly...", bin)
			}
		}

		if err := osutil.RunCommand(fmt.Sprintf(
			"docker-compose -f %s up --build", config.FilePath(workDir, config.DockerComposeFileName))); err != nil {
			zap.L().Fatal("Failed to start", zap.Error(err))
		}
	},
}
