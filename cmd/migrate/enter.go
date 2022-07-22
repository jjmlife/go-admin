package migrate

import "github.com/spf13/cobra"

var (
	config string
	mode string
	StartCmd = &cobra.Command{
		Use: "init",
		Short: "initialize the database",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

func init() {

}

func run() {
	
}