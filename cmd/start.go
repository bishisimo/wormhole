//@Time : 2020/11/6 下午2:29
//@Author : bishisimo
package cmd

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/exec"
)

var startCmd = &cobra.Command{
	Use:     "start",
	Aliases: []string{},
	Short:   "后台方式启动wormhole服务",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		f, err := os.OpenFile(viper.GetString("log_path"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			logrus.Errorln(err.Error())
		}
		logrus.SetOutput(f)
		exe := exec.Command(os.Args[0], "server")
		logPath := viper.GetString("log_path")
		if logPath != "" {
			stdout, err := os.OpenFile(logPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
			if err != nil {
				log.Println(os.Getpid(), ": 打开日志文件错误:", err)
			}
			exe.Stderr = stdout
			exe.Stdout = stdout
		}

		//异步启动子进程
		err = exe.Start()
		if err != nil {
			logrus.Fatalln("启动失败:", err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
