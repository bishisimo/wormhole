//@Time : 2020/11/3 上午10:35
//@Author : bishisimo
package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

var setCmd = &cobra.Command{
	Use:     "set",
	Aliases: []string{""},
	Short:   "设置环境参数",
	Long:    `use like: wormhole set self_name=new_wormhole`,
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			sp := strings.Split(arg, "=")
			if len(sp) != 2 {
				logrus.Errorln("参数格式无效,可能缺失符号\"=\"")
				return
			}
			if !viper.InConfig(sp[0]) {
				logrus.Errorln("不支持的参数:", sp[0])
			}
			viper.Set(sp[0], sp[1])
		}
		err := viper.WriteConfig()
		if err != nil {
			logrus.Errorln(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
