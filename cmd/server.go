package cmd

import (
	"github.com/spf13/cobra"
	"iads/server/internals/app/manager"
	//"iads/server"
	//"iads/server/model"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "restful api server",
	Run: func(cmd *cobra.Command, args []string) {
		//model.CreateTable()
		//server.ServerStart()
		manager.Run("")
	},
}
