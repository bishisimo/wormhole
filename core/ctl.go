// @time : 2021/5/31 15:50
// @author: bishisimo
// @describe: 服务器控制模块
package core

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
	"wormhole/event"
	"wormhole/protos/redux"
)

func (self *reduxServer) clear() {
	server.CurrentIndex = 0
	server.IndexMap = make(map[int]*redux.Device)
	server.HostMap = make(map[string]*redux.Device)
	server.NameMap = make(map[string]*redux.Device)
}

func SendSingle2Server(e int) {
	SingleChan <- e
}

//通过信号进行控制，可解耦外部命令为异步非阻塞控制
func Controller() {
	sysSingle := make(chan os.Signal, 1)
	signal.Notify(sysSingle, syscall.SIGINT, syscall.SIGTERM)
	server := GetReduxServer()
	for {
		select {
		case e := <-SingleChan:
			switch e {
			case event.Start:
				logrus.Infoln("start")
				go server.start()
			case event.Stop:
				logrus.Infoln("stop")
				server.stop()
				return
			case event.Restart:
				logrus.Infoln("restart")
				server.stop()
				go server.start()
			}
		case <-sysSingle:
			server.stop()
			return
		}
	}
}

func (self *reduxServer) start() {
	if viper.GetString("self_host") == "" {
		logrus.Fatalln("本机ip获取失败,无法启动服务,需手动添加到配置文件!")
	}

	port := viper.GetString("self_port")
	conn, _ := grpc.Dial("localhost:"+port, grpc.WithInsecure())
	c := redux.NewReduxClient(conn)
	_, err := c.CheckHealth(context.Background(), new(redux.Empty))
	if err == nil {
		logrus.Fatalln("运行失败:已有服务正在运行")
	}
	_ = conn.Close()

	self.Ctx, self.Cancel = context.WithCancel(context.Background())
	go UdpServer(self.Ctx)
	go Broadcast()
	nets := viper.GetStringSlice("nets")
	for _, targetNet := range nets {
		go func(targetNet string) {
			AddNet(targetNet)
		}(targetNet)
	}
	self.Listen()
}

func (self *reduxServer) stop() {
	logrus.Infoln("stop")
	if self.Cancel != nil {
		self.Cancel()
	}
	for host := range self.HostMap {
		logrus.Debug("stop and send stop event to host:", host)
		UdpHeat(host+":"+viper.GetString("udp_port"), event.OfflineEvent)
	}
	netMap := make(map[string]bool)
	historyNets := viper.GetStringSlice("nets")
	for _, host := range historyNets {
		netMap[host] = true
	}
	for host := range self.HostMap {
		if host == viper.GetString("self_host") {
			continue
		}
		netMap[host] = true
	}
	nets := make([]string, 0)
	for host := range netMap {
		nets = append(nets, host)
	}
	viper.Set("nets", nets)
	err := viper.WriteConfig()
	if err != nil {
		logrus.Errorln("写入配置文件失败", err.Error())
	}
	self.clear()
}

func (self *reduxServer) Stop(ctx context.Context, empty *redux.Empty) (*redux.Reply, error) {
	defer SendSingle2Server(event.Stop)
	return new(redux.Reply), nil
}

func (self *reduxServer) Restart(ctx context.Context, empty *redux.Empty) (*redux.Reply, error) {
	go SendSingle2Server(event.Restart)
	return new(redux.Reply), nil
}

func (self *reduxServer) Listen() {
	localName := viper.GetString("self_name")
	heat := &redux.Heat{
		Host: viper.GetString("self_host"),
		Name: localName,
		Port: viper.GetInt32("self_port"),
	}
	self.OnlineDevice(heat)
	lis, err := net.Listen("tcp", "0.0.0.0:"+viper.GetString("self_port"))
	if err != nil {
		logrus.Errorln(err.Error())
		return
	}
	s := grpc.NewServer()
	go func() {
		_ = <-self.Ctx.Done()
		s.Stop()
	}()
	redux.RegisterReduxServer(s, self)
	err = s.Serve(lis)
	if err != nil {
		logrus.Errorln("服务启动失败", err.Error())
	}
}
