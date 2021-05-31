//@Time : 2020/11/3 上午11:40
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

var stopCmd = &cobra.Command{
	Use:     "stop",
	Aliases: []string{""},
	Short:   "停止服务",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		port := viper.GetString("self_port")
		conn, err := grpc.Dial("localhost:"+port, grpc.WithInsecure())
		if err != nil {
			logrus.Errorln(err.Error())
			return
		}
		defer conn.Close()
		c := redux.NewReduxClient(conn)
		_, err = c.Stop(context.Background(), new(redux.Empty))
		if err != nil {
			logrus.Errorln("服务端未启动")
		}
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
