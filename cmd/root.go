package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	workDir string
	verbose bool

	rootCmd = &cobra.Command{
		Use:   "gebug",
		Short: "A tool for better debugging Go applications",
		Long: `Gebug helps you setup a fully suited debugging environment of Go application running inside container.
It enables options like connecting with remote debugger and put breakpoints inside the code or 
use hot-reload features which auto-build and run upon new change detected on the source code.`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			logConfig := zap.NewDevelopmentConfig()
			if !verbose {
				logConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
			}
			logger, err := logConfig.Build()
			if err != nil {
				return err
			}
			zap.ReplaceGlobals(logger)
			return nil
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			_ = zap.L().Sync()
			_ = zap.S().Sync()
		},
	}
)

// Execute executes the root cobra command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&workDir, "workdir", "w", ".", "your Go application root directory")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose mode")
}
