package cmd

import (
	"github.com/spf13/cobra"
	"snow-auth-service/migrations"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run all database migrations.",
	Long:  `This will run all database migrations that are needed--including creating the initial migrations table if it doesn't exist.'`,
	Run: func(cmd *cobra.Command, args []string) {
		migrations.MigrateDB()
	},
}
