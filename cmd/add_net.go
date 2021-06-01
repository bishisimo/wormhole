//@Time : 2020/10/14 上午11:18
//@Author : bishisimo
package cmd

import (
	"github.com/spf13/cobra"
	"regexp"
	"wormhole/core"
)

var addNewCmd = &cobra.Command{
	Use:     "net",
	Aliases: []string{"add", "an", "a"},
	Short:   "添加指定IP或IP所在网段(IP最后为255)",
	Long:    ``,
	Run: func(cmd *cobra.Command, args []string) {
		regIP := regexp.MustCompile(`\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3}`)
		if targetNet == "" && len(args) > 0 {
			targetNet = args[0]
		}
		if regIP.Match([]byte(targetNet)) {
			core.AddNet(targetNet)
		}
	},
}

var targetNet string

func init() {
	rootCmd.AddCommand(addNewCmd)
	addNewCmd.Flags().StringVarP(&targetNet, "net", "n", "", "指定添加网段")
}
