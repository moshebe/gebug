package cmd

import (
	"github.com/moshebe/gebug/pkg/setup"
	"github.com/pkg/errors"
	"os"
	"path"

	"github.com/hashicorp/go-multierror"
	"github.com/moshebe/gebug/pkg/config"
	"github.com/moshebe/gebug/pkg/osutil"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func init() {
	rootCmd.AddCommand(destroyCmd)
}

var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Destroy the Gebug project",
	Long:  "Remove all the Gebug related configuration files the project and perform cleanup",
	Run: func(cmd *cobra.Command, args []string) {
		var resultErr error
		zap.L().Debug("Cleaning project")
		if err := clean(); err != nil {
			resultErr = multierror.Append(resultErr, err)
		}
		zap.L().Debug("Deleting config directory")
		configDirPath := path.Join(workDir, config.RootDir)
		if osutil.FileExists(configDirPath) {
			if err := os.RemoveAll(configDirPath); err != nil {
				resultErr = multierror.Append(resultErr, err)
			}
		}

		zap.L().Debug("Disable Gebug configurations from detected IDEs")
		for name, ide := range setup.SupportedIde(workDir, 0) {
			detected, err := ide.Detected()
			if err != nil {
				resultErr = multierror.Append(resultErr, errors.WithMessagef(err, "detect IDE existence of '%s'", name))
				continue
			}

			if !detected {
				continue
			}

			err = ide.Disable()
			if err != nil {
				resultErr = multierror.Append(resultErr, errors.WithMessagef(err, "disable IDE '%s'", name))
			}
		}

		if resultErr != nil {
			zap.L().Fatal("Failed to destroy project", zap.Error(resultErr))
		}
	},
}
