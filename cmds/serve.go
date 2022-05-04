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
	Short: "Run the Snow Auth API Server.",
	Long:  `This will boot up the API Server that will allow Snow Auth to listen for http requests.`,
	Run: func(cmd *cobra.Command, args []string) {
		server.Start()
	},
}
