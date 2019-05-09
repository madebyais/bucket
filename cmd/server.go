package cmd

import (
	"log"

	"github.com/madebyais/bucket/app"
	"github.com/spf13/cobra"
)

var (
	port   string
	config string
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start bucket server",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		application := app.New(port, config)

		err := application.Start()
		if err != nil {
			log.Fatalf("Failed to start Bucket server, error=%s", err.Error())
		}
	},
}

func init() {
	serverCmd.PersistentFlags().StringVarP(&port, "port", "p", "8700", "Set bucket server port")
	serverCmd.PersistentFlags().StringVarP(&config, "config", "c", "bucket.yaml", "Set bucket config file path")
	mainCmd.AddCommand(serverCmd)
}
