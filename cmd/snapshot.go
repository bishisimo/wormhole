//@Time : 2020/11/16 下午5:31
//@Author : bishisimo
package cmd

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"wormhole/protos/redux"
)

var snapshotCmd = &cobra.Command{
	Use:     "snapshot",
	Aliases: []string{""},
	Short:   "设置环境参数",
	Long:    `use like: wormhole set self_name=new_wormhole`,
	Run: func(cmd *cobra.Command, args []string) {
		port := viper.GetString("self_port")
		conn, err := grpc.Dial("localhost:"+port, grpc.WithInsecure())
		if err != nil {
			logrus.Errorln(err.Error())
			return
		}
		defer conn.Close()
		c := redux.NewReduxClient(conn)
		_, err = c.Snapshot(context.Background(), new(redux.Empty))
		if err != nil {
			logrus.Errorln("服务端未启动")
		}
	},
}

func init() {
	rootCmd.AddCommand(snapshotCmd)
}
