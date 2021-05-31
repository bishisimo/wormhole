package cmd

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"wormhole/protos/redux"
)

var checkCmd = &cobra.Command{
	Use:     "check",
	Aliases: []string{},
	Short:   "检测本地服务端是否在线",
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
		res, err := c.CheckHealth(context.Background(), &redux.Empty{})
		if res == nil {
			fmt.Print(0)
		} else {
			fmt.Print(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
}
