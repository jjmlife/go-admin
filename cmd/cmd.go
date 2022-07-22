package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "go-admin",
	Short: "-v",
	SilenceUsage: true,
	DisableAutoGenTag: true,
	Long: `go-admin`,
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}


func init() {
	rootCmd.AddCommand() // app
	rootCmd.AddCommand() // migrate
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}