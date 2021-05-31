//@Time : 2020/11/6 下午5:12
//@Author : bishisimo
package cmd

import (
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{},
	Short:   "初始化配置文件",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		return
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
