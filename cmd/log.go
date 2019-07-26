package cmd

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
	"iads/lib/linux"
)

func init() {
	rootCmd.AddCommand(logCmd)
	logCmd.AddCommand(errCmd)
}

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "log functions",
}

var errCmd = &cobra.Command{
	Use:   "err",
	Short: "print err log",
	Run: func(cmd *cobra.Command, args []string) {
		msg := linux.NewMessages()
		err := msg.LoadEvent("/var/log/messages")
		if err != nil{
			log.Panic(err)
		}
		for i:=0; i<msg.EventList.Size();i++{
			fmt.Println(msg.EventList.Get(i))
		}
		
		/*lib.InitEnv()
		arr, err := lib.AnalysisLogFile("/var/log/messages")
		if err != nil {
			fmt.Println("/var/log/messages open filed.")
		} else {
			fmt.Println("/var/log/messages")
			if 0 == arr.Size() {
				fmt.Println("No errors.")
			} else {
				it := arr.Iterator()
				for it.Next() {
					tRow := it.Value().(*lib.RowLog)
					tStr := fmt.Sprintf("%d %s", tRow.Index, tRow.Data)
					fmt.Println(tStr)
				}
			}
		}

		arr, err = lib.AnalysisLogFile("/var/log/mcelog")
		if err != nil {
			fmt.Println("/var/log/mcelog open filed.")
		} else {
			fmt.Println("/var/log/mcelog")
			if 0 == arr.Size() {
				fmt.Println("No errors.")
			} else {
				it := arr.Iterator()
				for it.Next() {
					tRow := it.Value().(*lib.RowLog)
					tStr := fmt.Sprintf("%d %s", tRow.Index, tRow.Data)
					fmt.Println(tStr)
				}
			}
		}*/
	},
}
