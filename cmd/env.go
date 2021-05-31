//@Time : 2020/11/3 上午10:35
//@Author : bishisimo
package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"sort"
)

var envCmd = &cobra.Command{
	Use:     "env",
	Aliases: []string{""},
	Short:   "展示环境参数",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		keys := make([]string, 0)
		for _, key := range viper.AllKeys() {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			value := viper.GetString(key)
			if value != "" {
				fmt.Println(color.CyanString(key)+":", color.GreenString(value))
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(envCmd)
}
