//@Time : 2020/10/12 下午6:09
//@Author : bishisimo
package cmd

import (
	"github.com/spf13/cobra"
	"wormhole/core"
	"wormhole/event"
)

var serverCmd = &cobra.Command{
	Use:     "server",
	Aliases: []string{"run", "listen"},
	Short:   "启动wormhole服务",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		core.SendSingle2Server(event.Start)
		core.Controller()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
