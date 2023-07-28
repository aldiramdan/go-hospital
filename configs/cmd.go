package configs

import (
	"github.com/aldiramdan/hospital/databases/db"
	"github.com/spf13/cobra"
)

var initCommand = cobra.Command{
	Short: "backend golang",
	Long:  `backend golang with http native`,
}

func init() {
	initCommand.AddCommand(ServeCmd)
	initCommand.AddCommand(db.MigrateCmd)
}

func Run(args []string) error {
	initCommand.SetArgs(args)
	return initCommand.Execute()
}
