// @time : 2021/5/31 15:52
// @author: bishisimo
// @describe: 消息文件发送模块
package core

import (
	"bufio"
	"context"
	"errors"
	"github.com/cheggaaa/pb/v3"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"io"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"wormhole/protos/redux"
	"wormhole/utils"
)

func (self *reduxServer) destParseKey(dest string) *redux.DeviceKey {
	key := &redux.DeviceKey{
		Heat: &redux.Heat{
			Host:  "",
			Name:  "",
			Port:  0,
			Event: "",
		},
		Index: 0,
	}
	regHost := regexp.MustCompile(`\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3}:*\d{0,5}`) //ip或者ip:port
	regIndex := regexp.MustCompile(`\d+`)
	if regHost.Match([]byte(dest)) {
		sp := strings.Split(dest, ":")
		key.Heat.Host = sp[0]
		if len(sp) == 2 {
			port, err := strconv.Atoi(sp[1])
			if err == nil {
				key.Heat.Port = int32(port)
			}
		}
	} else {
		if regIndex.Match([]byte(dest)) {
			index, err := strconv.Atoi(dest)
			if err != nil {
				utils.HandleErr(err)
				return nil
			}
			key.Index = int32(index)
		} else {
			key.Heat.Name = dest
		}
	}
	return key
}

//发送文本消息
func (self *reduxServer) SendText(ctx context.Context, message *redux.Message) (*redux.Reply, error) {
	key := self.destParseKey(message.To)
	if ok := self.FillDeviceKey(key); ok {
		self.sendText(message.Data, key.Heat)
	} else {
		return nil, errors.New("FillDeviceKey fail")
	}
	return new(redux.Reply), nil
}

//发送文件
func (self *reduxServer) SendFile(ctx context.Context, message *redux.Message) (*redux.Reply, error) {
	key := self.destParseKey(message.To)
	if ok := self.FillDeviceKey(key); ok {
		self.sendFile(message.Data, key.Heat)
	} else {
		return nil, errors.New("FillDeviceKey fail")
	}
	return new(redux.Reply), nil
}

//自动检测文件或文本消息进行发送
func (self *reduxServer) SendAuto(ctx context.Context, message *redux.Message) (*redux.Reply, error) {
	if utils.IsExistPath(message.Data) {
		return self.SendFile(ctx, message)
	} else {
		return self.SendText(ctx, message)
	}
}

func (self *reduxServer) sendText(text string, heat *redux.Heat) {
	conn, err := grpc.Dial(heat.Host+":"+strconv.Itoa(int(heat.Port)), grpc.WithInsecure())
	if err != nil {
		log.Error().Stack().Err(err).Send()
		return
	}
	defer conn.Close()
	c := redux.NewReduxClient(conn)
	textMessage := &redux.TextMessage{
		FromHost: viper.GetString("self_host"),
		Text:     text,
	}
	_, err = c.AcceptText(context.Background(), textMessage)
	if err != nil {
		log.Error().Stack().Err(err).Send()
		return
	}
}

func (self *reduxServer) sendFile(targetFile string, heat *redux.Heat) {
	//根文件名称
	ads, err := filepath.Abs(targetFile)
	if err != nil {
		log.Error().Stack().Err(err).Send()
		return
	}
	ads = strings.ReplaceAll(ads, `\`, `/`)
	fileKey := path.Base(ads)
	//判断文件类型,添加子文件路径
	state, err := os.Stat(targetFile)
	if err != nil {
		utils.HandleErr(err)
		return
	}
	filePaths := make(map[string]string, 0)
	if state.IsDir() {
		filePaths, err = utils.GetAllFileSubPath(fileKey, targetFile)
		if err != nil {
			utils.HandleErr(err)
			return
		}
	} else {
		filePaths = map[string]string{targetFile: ""}
	}

	//连接rpc准备发送数据
	conn, err := grpc.Dial(heat.Host+":"+strconv.Itoa(int(heat.Port)), grpc.WithInsecure())
	if err != nil {
		utils.HandleErr(err)
		return
	}
	defer conn.Close()
	c := redux.NewReduxClient(conn)
	putStream, err := c.AcceptFile(context.Background())
	if putStream == nil {
		return
	}
	if err != nil {
		return
	}
	message := &redux.FileMessage{
		Key:      fileKey,
		FromHost: viper.GetString("self_host"),
	}
	var bar *pb.ProgressBar
	for realFilePath, softFilePath := range filePaths {
		message.Path = softFilePath
		if softFilePath == "" {
			message.Num = 0
			self.sendSingleFile(putStream, realFilePath, message, self.IsBacked)
		} else {
			if bar == nil {
				bar = pb.StartNew(len(filePaths))
			}
			message.Num = int32(len(filePaths))
			self.sendSingleFile(putStream, realFilePath, message, false)
			bar.Increment()
		}
	}
	if bar != nil {
		bar.Finish()
	}
	_, _ = putStream.CloseAndRecv()
}

//发送单文件数据,只处理文件数据,不处理消息元数据
func (self *reduxServer) sendSingleFile(putStream redux.Redux_AcceptFileClient, realFilePath string, message *redux.FileMessage, enableProgress bool) {
	var bar *pb.ProgressBar
	var barReader *pb.Reader
	state, err := os.Stat(realFilePath)
	if err != nil {
		utils.HandleErr(err)
		return
	}

	f, err := os.Open(realFilePath)
	if err != nil || f == nil {
		utils.HandleErr(err)
		return
	}

	defer f.Close()
	reader := bufio.NewReader(f)
	//填充要发送的元数据
	message.Size = int32(state.Size())
	message.Perm = int32(state.Mode().Perm())
	//对于单文件建立文件读取代理显示进度
	if enableProgress {
		bar = pb.Full.Start64(state.Size())
		barReader = bar.NewProxyReader(reader)
	}
	//发送数据

	for i := 0; ; i++ {
		n := 0
		buf := make([]byte, 4096)
		//从file读取到buf中
		if enableProgress {
			n, err = barReader.Read(buf)
		} else {
			n, err = reader.Read(buf)
		}
		if err != nil && err != io.EOF {
			utils.HandleErr(err)
		}
		//读取结束
		if n == 0 && i > 0 {
			if enableProgress {
				bar.Finish()
			}
			return
		}
		//流传输
		message.Data = buf[:n]
		err = putStream.Send(message)
		if err != nil {
			utils.HandleErr(err)
			return
		}
	}
}
