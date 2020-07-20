package cmd

import (
	"fmt"
	"gebug/pkg/config"
	"gebug/pkg/osutil"
	"github.com/hashicorp/go-multierror"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func init() {
	rootCmd.AddCommand(cleanCmd)
}

func clean() error {
	var resultErr error

	configFilePath := config.FilePath(workDir, config.DockerComposeFileName)
	if !osutil.FileExists(configFilePath) {
		if err := osutil.RunCommand(fmt.Sprintf(
			"docker-compose -f %s down -v --rmi all --remove-orphans", configFilePath)); err != nil {
			resultErr = multierror.Append(resultErr, err)
		}
	}

	if err := osutil.RunCommand("docker system prune -f"); err != nil {
		resultErr = multierror.Append(resultErr, err)
	}

	return resultErr
}

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean Gebug stack",
	Long:  "Cleanup docker-compose stack of the project",
	Run: func(cmd *cobra.Command, args []string) {
		err := clean()
		if err != nil {
			zap.L().Fatal("Failed to perform clean up", zap.Error(err))
		}
	},
}
