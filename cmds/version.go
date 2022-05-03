package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"snow-auth-service/configs"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Soul.",
	Long:  `All software has versions. This is Soul's.`,
	Run: func(cmd *cobra.Command, args []string) {
		version := fmt.Sprintf("Soul v%s", configs.GetVersion())
		fmt.Println(version)
	},
}
