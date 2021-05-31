//@Time : 2020/10/12 下午6:12
//@Author : bishisimo
package cmd

import (
	"context"
	"github.com/olekukonko/tablewriter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"os"
	"strconv"
	"wormhole/protos/redux"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "列出所有设备",
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
		res, err := c.ListDevice(context.Background(), &redux.Empty{})
		if res == nil || len(res.Devices) == 0 {
			logrus.Errorln("本地服务端离线!")
		} else {
			devices := res.Devices
			state := map[int32]string{
				0: "offline",
				1: "online",
			}
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"index", "Host", "Port", "Name", "State", "LastActiveTime"})
			for _, device := range devices {
				out := []string{
					strconv.Itoa(int(device.Key.Index)),
					device.Key.Heat.Host,
					strconv.Itoa(int(device.Key.Heat.Port)),
					device.Key.Heat.Name,
					state[device.StateCode],
					device.LastActiveTime,
				}
				rich := make([]tablewriter.Colors, len(out))
				if device.StateCode == 1 {
					rich[4] = tablewriter.Colors{tablewriter.Normal, tablewriter.FgCyanColor}
				} else {
					rich[4] = tablewriter.Colors{tablewriter.Normal, tablewriter.FgHiRedColor}
				}
				table.Rich(out, rich)
			}
			table.Render()
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
