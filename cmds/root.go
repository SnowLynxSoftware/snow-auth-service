package cmd

// CONFIG: https://github.com/spf13/cobra/blob/master/user_guide.md

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "soul",
	Short: "Soul is a CLI for the Soul Arenas Game.",
	Long: `Soul is a CLI for the Soul Arenas Game  
built with love by Dylan Legendre!`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("For help, type: soul help")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
