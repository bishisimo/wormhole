// @time : 2021/6/1 17:54
// @author: bishisimo
// @describe:
package cmd

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"wormhole/protos/redux"
	"wormhole/task"
)

var sendThinCmd = &cobra.Command{
	Use:     "send",
	Aliases: []string{"s"},
	Short:   "发送文本消息或文件",
	Long: `
1、wormhole send 0 "this is message"
2、wormhole send 192.168.1.108 /path/for/file
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			logrus.Info("无效的参数")
			_ = cmd.Help()
		} else if len(args) >= 2 {
			c := task.NewLocalClient()
			if c == nil {
				logrus.Errorln("本地服务端离线！")
				return
			}
			defer c.Close()
			for i := 1; i < len(args); i++ {
				msg := &redux.Message{
					Data: args[i],
					From: "",
					To:   args[0],
				}
				_, err := c.SendAuto(context.Background(), msg)
				if err != nil {
					logrus.Errorln(err)
					return
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(sendThinCmd)
}
