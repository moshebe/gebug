package cmd

import (
	"fmt"

	"github.com/moshebe/gebug/version"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Gebug version",
	Long:  "Print the version number of Gebug",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.Name())
	},
}
