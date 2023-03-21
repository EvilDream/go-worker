package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of application",
	Long:  `Print the running version of application go-worker`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You're running go-worker with the following version:\n v0.0.1")
	},
}
