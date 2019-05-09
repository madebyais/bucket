package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var mainCmd = &cobra.Command{
	Use:   "bucket",
	Short: "Bucket is a file server storage using REST",
}

// Execute is used to execute main command for bucket
func Execute() {
	err := mainCmd.Execute()
	if err != nil {
		fmt.Printf(`Error in executing command for bucket, error=%s`, err.Error())
		os.Exit(1)
	}
}
