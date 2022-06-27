// Copyright © 2022 hu-jinwen, hu-jinwen@outlook.com

package cmd

import (
	"fmt"
	_ "github.com/hu-jinwen/Typora-img-upload/cmd/sub"
	"github.com/spf13/cobra"
	"os"
)

var (
	cfgFile string
)

func init() {
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is ./.Typora-img-upload.yaml)")
}

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "Typora-img-upload",
	Short: "Typora 图片上传至图床插件",
	Long:  `Typora 图片自动上传图床插件`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		fmt.Println("请输入 command，例如 Typora-img-upload oss。输入 -h 查看帮助。")
		return err
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
