package cmd

import (
	"github.com/spf13/cobra"
	"snow-auth-service/internal/server"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run the Soul API Server.",
	Long:  `This will boot up the API Server that will allow Soul to listen for http requests.`,
	Run: func(cmd *cobra.Command, args []string) {
		server.Start()
	},
}
