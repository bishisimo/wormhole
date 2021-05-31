package task

import (
	"bufio"
	"context"
	"github.com/cheggaaa/pb/v3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"io"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"wormhole/protos/redux"
	"wormhole/utils"
)

func SendText(text string, heat *redux.Heat) {
	conn, err := grpc.Dial(heat.Host+":"+strconv.Itoa(int(heat.Port)), grpc.WithInsecure())
	if err != nil {
		logrus.Errorln(err.Error())
		return
	}
	defer conn.Close()
	c := redux.NewReduxClient(conn)
	textMessage := &redux.TextMessage{
		FromHost: viper.GetString("self_host"),
		Text:     text,
	}
	_, err = c.SendText(context.Background(), textMessage)
	if err != nil {
		logrus.Errorln(err.Error())
		return
	}
}
func SendFile(targetFile string, heat *redux.Heat) {
	//根文件名称
	ads, err := filepath.Abs(targetFile)
	if err != nil {
		logrus.Errorln(err.Error())
	}
	ads = strings.ReplaceAll(ads, `\`, `/`)
	fileKey := path.Base(ads)
	//判断文件类型,添加子文件路径
	state, err := os.Stat(targetFile)
	if err != nil {
		logrus.Errorln(err.Error())
		return
	}
	filePaths := make(map[string]string, 0)
	if state.IsDir() {
		filePaths, err = utils.GetAllFileSubPath(fileKey, targetFile)
		if err != nil {
			logrus.Errorln(err.Error())
			return
		}
	} else {
		filePaths = map[string]string{targetFile: ""}
	}

	//连接rpc准备发送数据
	conn, err := grpc.Dial(heat.Host+":"+strconv.Itoa(int(heat.Port)), grpc.WithInsecure())
	if err != nil {
		logrus.Errorln(err)
		return
	}
	defer conn.Close()
	c := redux.NewReduxClient(conn)
	putStream, err := c.SendFile(context.Background())
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
			SendSingleFile(putStream, realFilePath, message, true)
		} else {
			if bar == nil {
				bar = pb.StartNew(len(filePaths))
			}
			message.Num = int32(len(filePaths))
			SendSingleFile(putStream, realFilePath, message, false)
			bar.Increment()
		}
	}
	if bar != nil {
		bar.Finish()
	}
	_, _ = putStream.CloseAndRecv()
}

//发送单文件数据,只处理文件数据,不处理消息元数据
func SendSingleFile(putStream redux.Redux_SendFileClient, realFilePath string, message *redux.FileMessage, enableProgress bool) {
	var bar *pb.ProgressBar
	var barReader *pb.Reader
	state, err := os.Stat(realFilePath)
	if err != nil {
		logrus.Errorln(err.Error())
		return
	}

	f, err := os.Open(realFilePath)
	if err != nil {
		logrus.Errorln(err.Error())
		return
	}
	if f == nil {
		logrus.Errorln(realFilePath + "无法打开")
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
			logrus.Errorln(err.Error())
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
			logrus.Errorln(err.Error())
			return
		}
	}
}
