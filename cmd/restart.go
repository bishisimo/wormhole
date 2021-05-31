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

var restartCmd = &cobra.Command{
	Use:     "restart",
	Aliases: []string{""},
	Short:   "重启服务",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		port := viper.GetString("self_port")
		conn, err := grpc.Dial("localhost:"+port, grpc.WithInsecure())
		if err != nil {
			logrus.Error(err.Error())
			return
		}
		defer conn.Close()
		c := redux.NewReduxClient(conn)
		_, err = c.Restart(context.Background(), new(redux.Empty))
		if err != nil {
			logrus.Error(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(restartCmd)
}
