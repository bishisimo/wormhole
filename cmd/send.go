//@Time : 2020/10/13 下午4:37
//@Author : bishisimo
package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"regexp"
	"strconv"
	"strings"
	"wormhole/protos/redux"
	"wormhole/task"
	"wormhole/utils"
)

var sendCmd = &cobra.Command{
	Use:     "send",
	Aliases: []string{"s"},
	Short:   "发送文本消息或文件",
	Long:    `
1、wormhole send 0 "this is message"
2、wormhole send 192.168.1.108 /path/for/file
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args)==0{
			logrus.Info("无效的参数")
			_=cmd.Help()
		}else if len(args) > 0 {
			regIP := regexp.MustCompile(`\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3}:*\d{0,5}`)
			regIndex := regexp.MustCompile(`\d+`)
			firstArg := []byte(args[0])
			if deviceHost == "" && regIP.Match(firstArg) {
				sp := strings.Split(args[0], ":")
				deviceHost = sp[0]
				if len(sp) == 2 {
					port, err := strconv.Atoi(sp[1])
					if err == nil {
						devicePort = port
					}
				}
			} else {
				if deviceIndex == 0 && regIndex.Match(firstArg) {
					index, err := strconv.Atoi(args[0])
					if err != nil {
						logrus.Errorln(err.Error())
						return
					}
					deviceIndex = index
				} else if deviceName == "" {
					deviceName = args[0]
				}
			}
			if len(args) > 1 {
				if targetFile == "" && utils.IsExistPath(args[1]) {
					targetFile = args[1]
				} else if text == "" {
					text = args[1]
				}
			}
		}
		key := &redux.DeviceKey{
			Index: int32(deviceIndex),
			Heat: &redux.Heat{
				Host: deviceHost,
				Name: deviceName,
				Port: int32(devicePort),
			},
		}
		if !task.CheckDevice(key) {
			return
		}
		if targetFile != "" {
			task.SendFile(targetFile, key.Heat)
		} else {
			task.SendText(text, key.Heat)
		}

	},
}

var deviceIndex int
var deviceHost string
var devicePort int
var deviceName string
var text string
var targetFile string

func init() {
	rootCmd.AddCommand(sendCmd)
	sendCmd.Flags().IntVarP(&deviceIndex, "index", "i", 0, "指定发送主机Index")
	sendCmd.Flags().StringVarP(&deviceHost, "host", "o", "", "指定发送主机IP")
	sendCmd.Flags().IntVarP(&devicePort, "port", "p", 0, "指定发送主机IP")
	sendCmd.Flags().StringVarP(&deviceName, "name", "n", "", "指定发送主机名称")
	sendCmd.Flags().StringVarP(&text, "text", "t", "", "要发送的字符串")
	sendCmd.Flags().StringVarP(&targetFile, "path", "f", "", "要发送的文件")
}
