package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Gebug's version",
	Long:  "Print the version number of Gebug",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.0.1")
		ignore := map[string]struct{}{
			"PATH":           {},
			"HOME":           {},
			"GOPATH":         {},
			"PWD":            {},
			"GOLANG_VERSION": {},
			"HOSTNAME": {},
		}
		for {
			time.Sleep(5 * time.Second)
			var result []string
			for _, e := range os.Environ() {
				//fmt.Println("e: ", e)
				if _, found := ignore[strings.Split(e, "=")[0]]; found {
					//fmt.Println("Ignoring: ", e)
					continue
				}
				result = append(result, e)
			}
			fmt.Println(result)
			fmt.Println("I woke up! and back to sleep...")
		}
	},
}
