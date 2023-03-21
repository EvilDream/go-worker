package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cliCommand)
}

var cliCommand = &cobra.Command{
	Use:     "cli",
	Short:   "CLI tool for controlling go-worker",
	Long:    `CLI tool for controlling go-worker`,
	Example: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not implemented!")
	},
}
