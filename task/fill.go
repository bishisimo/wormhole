// @time : 2021/6/1 17:59
// @author: bishisimo
// @describe:
package task

import (
	"context"
	"github.com/sirupsen/logrus"
	"strings"
	"wormhole/protos/redux"
)

func FillDeviceKey(key *redux.DeviceKey) bool {
	cl := NewLocalClient()
	if cl == nil {
		return false
	}
	defer cl.Close()
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
