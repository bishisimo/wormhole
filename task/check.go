package task

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"strings"
	"wormhole/protos/redux"
)

func CheckDevice(key *redux.DeviceKey) bool {
	if key.Heat.Host != "" && key.Heat.Port != 0 {
		return true
	}
	conn_local, err := grpc.Dial("localhost:"+viper.GetString("self_port"), grpc.WithInsecure())
	if err != nil || conn_local == nil {
		logrus.Errorln("本地服务端离线!")
		return false
	}
	defer conn_local.Close()
	cl := redux.NewReduxClient(conn_local)

	device, err := cl.GetDevice(context.Background(), key)
	if err != nil {
		if strings.Contains(err.Error(), "Unavailable") {
			logrus.Errorln("本地服务端离线!")
		} else {
			logrus.Errorln("无法找到指定设备!")
		}
		return false
	}
	if device.StateCode == 0 {
		logrus.Errorln("指定设备离线,无法接收!")
		return false
	}
	key.Heat = device.Key.Heat
	return true
}
