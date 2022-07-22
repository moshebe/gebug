package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Tag    string
	Commit string
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Gebug version",
	Long:  "Print the version number of Gebug",
	Run: func(cmd *cobra.Command, args []string) {
		var v string
		switch {
		case Tag != "":
			v = Tag
		case Commit != "":
			v = "dev-" + Commit
		default:
			v = "development"
		}

		fmt.Println("version: ", v)
	},
}
