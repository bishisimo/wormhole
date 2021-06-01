// @time : 2021/5/31 15:49
// @author: bishisimo
// @describe: 服务器接收消息文件模块
package core

import (
	"bufio"
	"context"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/fatih/color"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
	"wormhole/protos/redux"
	"wormhole/utils"
)

//接收消息
func (self *reduxServer) AcceptText(ctx context.Context, message *redux.TextMessage) (*redux.Reply, error) {
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

//发送文件消息
func (self *reduxServer) AcceptFile(server redux.Redux_AcceptFileServer) error {
	defer func() {
		if err := recover(); err != nil {
			utils.HandleErr(err.(error))
		}
	}()
	//生成储存文件夹
	acceptDir := viper.GetString("accept_dir")
	if !utils.IsExistPath(acceptDir) {
		_ = os.MkdirAll(acceptDir, 0777)
	}
	var fileBar *pb.ProgressBar
	var barWriter *pb.Writer
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
			log.Error().Msg("解析文件出错")
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
					utils.HandleErr(err)
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
				utils.HandleErr(err)
			}
			defer f.Close()
			writer = bufio.NewWriter(f)
			if message.Num == 0 && self.IsBacked {
				fileBar = pb.Full.Start64(int64(message.Size))
				//defer fileBar.Finish()
				barWriter = fileBar.NewProxyWriter(writer)

			} else if self.IsBacked {
				if dirBar == nil {
					dirBar = pb.StartNew(int(message.Num))
					//defer dirBar.Finish()
				}
				dirBar.Increment()
			}
		}
		//写入数据
		if message.Num == 0 {
			if barWriter != nil {
				_, err = barWriter.Write(message.Data)
			} else {
				_, err = writer.Write(message.Data)
			}
			err := writer.Flush()
			if err != nil {
				utils.HandleErr(err)
				return err
			}
		} else {
			_, err = writer.Write(message.Data)
			err := writer.Flush()
			if err != nil {
				utils.HandleErr(err)
				return err
			}
		}
		if err != nil {
			utils.HandleErr(err)
		}
		state, err := os.Stat(f.Name())
		if err != nil {
			utils.HandleErr(err)
			return err
		}
		if fileBar != nil && message.Num == 0 && state.Size() == int64(message.Size) {
			fileBar.Finish()
		} else if dirBar != nil && message.Num > 0 && dirBar.Current() == dirBar.Total() {
			dirBar.Finish()
		}
	}
}
