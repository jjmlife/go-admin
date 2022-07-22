package app

import "github.com/spf13/cobra"

var (
	config string
	mode string
	StartCmd = &cobra.Command{
		Use: "server",
		Short: "start api server",
		Example: "go-admin server config/setting.yml",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {

}

func run() {
	
}