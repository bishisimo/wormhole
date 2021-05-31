//@Time : 2020/10/14 上午9:11
//@Author : bishisimo
package core

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net"
	"strconv"
	"strings"
	"sync"
	"wormhole/event"
	"wormhole/protos/redux"
)

func UdpServer(ctx context.Context) {
	//创建监听的地址，并且指定udp协议
	udp_addr, err := net.ResolveUDPAddr("udp", "0.0.0.0:"+viper.GetString("udp_port"))
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	conn, err := net.ListenUDP("udp", udp_addr) //创建数据通信socket
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	go func() {
		_ = <-ctx.Done()
		err := conn.Close()
		if err != nil {
			logrus.Error(err.Error())
		}
	}()
	defer conn.Close()
	logrus.Debug("udp server standby!")
	for {
		buf := make([]byte, 1024)
		n, _, err := conn.ReadFromUDP(buf) //接收客户端发送过来的数据，填充到切片buf中。
		if err != nil {
			return
		}
		udpMessage := new(UdpMessage)
		err = json.Unmarshal(buf[:n], udpMessage)
		if err != nil {
			logrus.Error(err.Error())
			return
		}
		if udpMessage.Heat.Host == viper.GetString("self_host") {
			logrus.Debug("udp server accept self message")
			continue
		}
		reduxServer := GetReduxServer()
		logrus.Debugf("%+v\n", *udpMessage)
		switch udpMessage.Heat.Event {
		case event.OfflineEvent:
			reduxServer.EventAdd(udpMessage.Heat)
		case event.RespondEvent:
			reduxServer.EventAdd(udpMessage.Heat)
		case event.RepeatEvent:
			logrus.Warn("当前网络有重名,建议修改名称")
		case event.BroadcastEvent:
			{
				device := reduxServer.GetDeviceByName(udpMessage.Heat.Name)
				if device != nil && device.Key.Heat.Host != udpMessage.Heat.Host { //与已知设备冲突
					logrus.Warn("检测到"+device.Key.Heat.Host, "与", udpMessage.Heat.Host, "有重名:", udpMessage.Heat.Name)
					UdpHeat(udpMessage.Heat.Host+":"+strconv.Itoa(udpMessage.Port), event.RepeatEvent)
				} else {
					UdpHeat(udpMessage.Heat.Host+":"+strconv.Itoa(udpMessage.Port), event.RespondEvent)
					reduxServer.EventAdd(udpMessage.Heat)
				}
			}
		default:
			logrus.Warn("遇到未知类型:", udpMessage.Heat.Event)
		}
	}
}

func UdpHeat(address string, heatType string) {
	conn, err := net.Dial("udp", address)
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	defer conn.Close()
	message := UdpMessage{
		Heat: &redux.Heat{
			Host:  viper.GetString("self_host"),
			Name:  viper.GetString("self_name"),
			Port:  int32(viper.GetInt("self_port")),
			Event: heatType,
		},
		Port: 10808,
	}
	data, err := json.Marshal(message)
	conn.Write(data)
}

func Broadcast() {
	UdpHeat("255.255.255.255:"+viper.GetString("udp_port"), event.BroadcastEvent)
}

func AddNet(targetNet string) {
	sub := strings.Split(targetNet, ".")
	if sub[len(sub)-1] == "255" {
		head := sub[0] + "." + sub[1] + "." + sub[2] + "."
		wg := sync.WaitGroup{}
		for i := 1; i < 255; i++ {
			wg.Add(1)
			go func(p int) {
				UdpHeat(head+strconv.Itoa(p)+":"+viper.GetString("udp_port"), event.BroadcastEvent)
				wg.Done()
			}(i)
		}
		wg.Wait()
	} else {
		UdpHeat(targetNet+":"+viper.GetString("udp_port"), event.BroadcastEvent)
	}
}

type UdpMessage struct {
	Heat *redux.Heat `json:"key"`
	Port int         `json:"port"`
}
