package cmd

// CONFIG: https://github.com/spf13/cobra/blob/master/user_guide.md

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "snow-auth-service",
	Short: "Snow Auth Service is a CLI for SnowLynxSoftware Authentication Microservice.",
	Long: `Snow Auth Service is a CLI for SnowLynxSoftware Authentication Microservice  
built with love by Dylan Legendre!`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("For help, type: snow-auth-service help")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
