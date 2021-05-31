//@Time : 2020/10/12 上午11:27
//@Author : bishisimo
package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/exec"
	"path"
	"strings"
	"wormhole/cmd"
	"wormhole/utils"
)

func main() {
	sourceDir, err := exec.LookPath(os.Args[0])
	if err != nil {
		logrus.Errorln("无法检测程序目录", err.Error())
	}
	sourceDir = strings.ReplaceAll(sourceDir, `\`, `/`)
	baseDir := path.Dir(sourceDir)
	viper.SetDefault("self_port", 1808)
	viper.SetDefault("udp_port", 10808)
	viper.SetDefault("accept_dir", path.Join(baseDir, "accept"))
	viper.SetDefault("self_name", utils.NcInfo.Mac)
	viper.SetDefault("log_lever", "info")
	viper.SetDefault("ip_mode", "auto")
	viper.SetDefault("log_lever", "info")

	configDir := baseDir
	configFile := "wormhole_config.yaml"
	logPath := path.Join(configDir, "wormhole.log")
	viper.SetDefault("log_path", logPath)

	configPath := path.Join(configDir, configFile)
	viper.SetConfigName(strings.Split(configFile, ".")[0])
	viper.SetConfigType(strings.Split(configFile, ".")[1])
	viper.AddConfigPath(configDir)

	if !utils.IsExistPath(configPath) {
		err := os.MkdirAll(configDir, 0777)
		err = viper.SafeWriteConfigAs(configPath)
		if err != nil {
			logrus.Error("初始化配置失败", err.Error())
		}
	}
	err = viper.ReadInConfig()
	if err != nil {
		logrus.Error("读取配置文件失败!", err.Error())
	}
	lever, err := logrus.ParseLevel(viper.GetString("log_lever"))

	if err == nil {
		logrus.SetLevel(lever)
	}
	logrus.Debugln("log-lever:", viper.GetString("log_lever"))

	if viper.GetString("self_host") == "" || viper.GetString("ip_mode") == "auto" {
		ip := utils.NcInfo.Intranet
		viper.Set("self_host", ip)
		logrus.Debugln("动态获取IP:", ip)
	}

	viper.WatchConfig()
	cmd.Execute()
}
