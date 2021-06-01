// @time : 2021/5/31 15:50
// @author: bishisimo
// @describe: 设备管理模块
package core

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"sort"
	"time"
	"wormhole/protos/redux"
)

func (self *reduxServer) GetDeviceByHost(host string) *redux.Device {
	if device, ok := self.HostMap[host]; ok {
		return device
	} else {
		return nil
	}
}

//通过名称获取服务端管理的设备信息
func (self *reduxServer) GetDeviceByName(name string) *redux.Device {
	if device, ok := self.NameMap[name]; ok {
		return device
	} else {
		return nil
	}
}
func (self *reduxServer) GetDevice(ctx context.Context, key *redux.DeviceKey) (*redux.Device, error) {
	logrus.Traceln("%+v\n", key)
	if key.Heat.Name != "" {
		deviceP, ok := self.NameMap[key.Heat.Name]
		if ok {
			logrus.Debugf("%+v\n", deviceP)
			return deviceP, nil
		}
	} else if key.Heat.Host != "" {
		deviceP, ok := self.HostMap[key.Heat.Host]
		if ok {
			logrus.Debugf("%+v\n", deviceP)
			return deviceP, nil
		}
	} else {
		deviceP, ok := self.IndexMap[int(key.Index)]
		if ok {
			logrus.Debugf("%+v\n", deviceP)
			return deviceP, nil
		}
	}
	logrus.Debugln("cannot find device")
	return nil, errors.New("cannot find device")
}

func (self *reduxServer) OnlineDevice(heat *redux.Heat) {
	if deviceP, ok := self.HostMap[heat.Host]; ok {
		deviceP.Key.Heat = heat
		deviceP.StateCode = 1
		deviceP.LastActiveTime = time.Now().Format("2006-01-02 15:04:05")
	} else {
		key := &redux.DeviceKey{
			Heat:  heat,
			Index: int32(self.CurrentIndex),
		}
		deviceP = &redux.Device{
			Key:            key,
			StateCode:      1,
			LastActiveTime: time.Now().Format("2006-01-02 15:04:05"),
		}
		self.IndexMap[self.CurrentIndex] = deviceP
		self.CurrentIndex++
		self.NameMap[heat.Name] = deviceP
		self.HostMap[heat.Host] = deviceP
	}
}

func (self *reduxServer) OfflineDevice(heat *redux.Heat) {
	if deviceP, ok := self.HostMap[heat.Host]; ok {
		deviceP.StateCode = 0
	}
}

func (self *reduxServer) BeatHeat(ctx context.Context, heat *redux.Heat) (*redux.Reply, error) {
	self.OnlineDevice(heat)
	return &redux.Reply{
		Status: 0,
	}, nil
}

func (self *reduxServer) CheckHealth(ctx context.Context, empty *redux.Empty) (*redux.Reply, error) {
	return new(redux.Reply), nil
}

func (self *reduxServer) ListDevice(ctx context.Context, e *redux.Empty) (*redux.ListResponse, error) {
	devices := make([]*redux.Device, 0)
	for _, v := range self.HostMap {
		devices = append(devices, v)
	}
	sort.Slice(devices, func(i, j int) bool {
		return devices[i].Key.Index < devices[j].Key.Index
	})
	response := &redux.ListResponse{
		Devices: devices,
	}
	return response, nil
}

func (self *reduxServer) FillDeviceKey(key *redux.DeviceKey) bool {
	if key.Heat.Host != "" && key.Heat.Port != 0 {
		return true
	}

	device, err := self.GetDevice(context.Background(), key)
	if err != nil {
		logrus.Errorln("无法找到指定设备!")
		return false
	}
	if device.StateCode == 0 {
		logrus.Errorln("指定设备离线,无法接收!")
		return false
	}
	key.Heat = device.Key.Heat
	return true
}
