//@Time : 2020/10/12 下午6:00
//@Author : bishisimo
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "wormhole",
	Short: "wormhole是跨平台数据传输软件,支持局域网自动发现",
	Long: `使用示例: 
1、启动本地服务：wormhole server 
2、查看在线设备：wormhole ls
3、发送消息：wormhole send 0 hello
4、发送文件：wormhole send ./wormhole
`,
	Run: func(cmd *cobra.Command, args []string) {
		_=cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("cmd error")
		os.Exit(1)
	}
}