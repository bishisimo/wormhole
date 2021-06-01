//@Time : 2020/10/13 上午10:12
//@Author : bishisimo
package core

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/bradleyjkemp/memviz"
	"github.com/cheggaaa/pb/v3"
	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"io"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"sort"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"
	"wormhole/event"
	"wormhole/protos/redux"
	"wormhole/utils"
)

type reduxServer struct {
	CurrentIndex int
	IndexMap     map[int]*redux.Device
	HostMap      map[string]*redux.Device
	NameMap      map[string]*redux.Device
	EventChan    chan *redux.Heat
	Ctx          context.Context
	Cancel       context.CancelFunc
	MessageChan  chan string
}

//提供本地消息流，给外部程序使用
func (self *reduxServer) GetMessageStream(empty *redux.Empty, streamServer redux.Redux_GetMessageStreamServer) error {
	for message := range self.MessageChan {
		err := streamServer.Send(&redux.Message{Data: message})
		if err != nil {
			return err
		}
	}
	return nil
}

var server *reduxServer
var once sync.Once
var SingleChan chan int

func init() {
	SingleChan = make(chan int, 1)
}

func GetReduxServer() *reduxServer {
	once.Do(func() {
		server = &reduxServer{
			CurrentIndex: 0,
			IndexMap:     make(map[int]*redux.Device),
			HostMap:      make(map[string]*redux.Device),
			NameMap:      make(map[string]*redux.Device),
			EventChan:    make(chan *redux.Heat, 255),
			MessageChan:  make(chan string, 255),
		}
		go server.EventHandle()
	})
	return server
}

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

func (self *reduxServer) EventAdd(heat *redux.Heat) {
	self.EventChan <- heat
}
func (self *reduxServer) EventHandle() {
	for heat := range self.EventChan {
		if heat.Event == event.OfflineEvent {
			self.OfflineDevice(heat)
		} else {
			self.OnlineDevice(heat)
		}
	}
}

func (self *reduxServer) SaveMessage(message string) {
	if len(self.MessageChan) > 250 {
		<-self.MessageChan
	}
	self.MessageChan <- message
}

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

func (self *reduxServer) SendText(ctx context.Context, message *redux.TextMessage) (*redux.Reply, error) {
	fmt.Println(
		color.CyanString("["+time.Now().Format("2006-01-02 15:04:05")+"]"),
		color.YellowString(message.FromHost)+color.MagentaString("'Message:"),
		color.BlueString(message.Text),
	)
	self.SaveMessage("[" + time.Now().Format("2006-01-02 15:04:05") + "]" + message.FromHost + "'Message:\n" + message.Text)
	return &redux.Reply{
		Status: 0,
	}, nil
}

func (self *reduxServer) SendFile(server redux.Redux_SendFileServer) error {
	defer func() {
		recover()
	}()
	//生成储存文件夹
	acceptDir := viper.GetString("accept_dir")
	if !utils.IsExistPath(acceptDir) {
		_ = os.MkdirAll(acceptDir, 0777)
	}
	var fileBar *pb.ProgressBar
	var barReader *pb.Writer
	var dirBar *pb.ProgressBar

	var f *os.File
	fileKey := ""
	filePath := ""
	isFirstMessage := true
	var writer *bufio.Writer
	for {
		message, err := server.Recv()
		//r全部数据接收完毕
		if err == io.EOF {
			return nil
		}
		if err != nil {
			logrus.Errorln("解析文件出错")
			return err
		}
		//根文件名处理
		if isFirstMessage {
			isFirstMessage = false
			//生成文件或文件夹
			fileKey = message.Key
			if utils.IsExistPath(path.Join(acceptDir, message.Key)) {
				fileNameSp := strings.Split(message.Key, ".")
				for i, s := range fileNameSp {
					if i == 0 {
						fileKey = s + "_" + strconv.FormatInt(time.Now().Unix(), 10)
					} else {
						fileKey += "." + s
					}
				}
			}
			var keyString string
			if message.Num == 0 {
				keyString = color.GreenString(message.Key)
			} else {
				keyString = color.BlueString(message.Key)
			}
			var fileTypeStr string
			if message.Num == 0 {
				fileTypeStr = "'file:"
			} else {
				fileTypeStr = "'dir:"
			}
			fmt.Println(
				color.CyanString("["+time.Now().Format("2006-01-02 15:04:05")+"]"),
				color.YellowString(message.FromHost)+color.MagentaString(fileTypeStr),
				keyString,
			)
			self.SaveMessage("[" + time.Now().Format("2006-01-02 15:04:05") + "]" + message.FromHost + fileTypeStr + "\n" + keyString)
		}

		//文件句柄
		if message.Num > 0 && filePath != message.Path || message.Num == 0 && filePath == "" {
			//记录当前文件路径
			filePath = message.Path
			//创建必要局部文件夹
			if filePath != "" {
				dir := path.Join(acceptDir, fileKey, path.Dir(filePath))
				err = utils.MakeDir(dir)
				if err != nil {
					logrus.Errorln("创建文件夹失败", err.Error())
				}
			} else {
				filePath = fileKey
			}
			//dir:=path.Join(acceptDir,path.Dir(filePath))
			//打开文件句柄
			if f != nil {
				f.Close()
			}
			f, err = os.OpenFile(path.Join(acceptDir, fileKey, message.Path), os.O_WRONLY|os.O_CREATE, os.FileMode(message.Perm))
			if err != nil {
				logrus.Errorln("写入文件失败", err.Error())
			}
			defer f.Close()
			writer = bufio.NewWriter(f)
			if message.Num == 0 {
				fileBar = pb.Full.Start64(int64(message.Size))
				//defer fileBar.Finish()
				barReader = fileBar.NewProxyWriter(writer)

			} else {
				if dirBar == nil {
					dirBar = pb.StartNew(int(message.Num))
					//defer dirBar.Finish()
				}
				dirBar.Increment()
			}
		}
		//写入数据
		if message.Num == 0 {
			_, err = barReader.Write(message.Data)
			err := writer.Flush()
			if err != nil {
				logrus.Errorln(err.Error())
				return err
			}
		} else {
			_, err = writer.Write(message.Data)
			err := writer.Flush()
			if err != nil {
				logrus.Errorln(err.Error())
				return err
			}
		}
		if err != nil {
			logrus.Errorln(err.Error())
		}
		state, err := os.Stat(f.Name())
		if err != nil {
			logrus.Errorln(err.Error())
			return err
		}
		if message.Num == 0 && state.Size() == int64(message.Size) {
			fileBar.Finish()
		} else if message.Num > 0 && dirBar.Current() == dirBar.Total() {
			dirBar.Finish()
		}
	}
}

func (self *reduxServer) Snapshot(ctx context.Context, empty *redux.Empty) (*redux.Reply, error) {
	f, err := os.OpenFile("./core.dot", os.O_WRONLY|os.O_CREATE, os.FileMode(0777))
	if err != nil {
		logrus.Fatalln(err.Error())
	}
	memviz.Map(f, self)
	command := exec.Command("dot", "-Tpng", "core.dot", "-o", "core.png")
	err = command.Run()
	if err != nil {
		logrus.Warningln(err.Error())
	}
	return new(redux.Reply), nil
}
