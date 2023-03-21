package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-worker",
	Short: "go-worker is a service for processing jobs",
	Long: `Go-worker is a service and CLI that used for working with any kind of jobs.
Service is a processor that processes jobs that are coming into go-worker.
CLI is a tool for controlling go-worker.
Jobs are expandable.
Complete documentation is available at https://github.com/evildream/README.md`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
