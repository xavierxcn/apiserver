/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xavierxcn/apiserver/pkg/log"
	"github.com/xavierxcn/apiserver/pkg/token"
)

// tokenCmd represents the token command
var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "创建一个token",
	Long:  `创建一个token`,
	Run: func(cmd *cobra.Command, args []string) {
		tokenString, err := token.Sign(nil, token.Context{Name: tokenName}, viper.GetString("serve.jwt_secret"))
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(tokenString)
	},
}

var (
	tokenName string
)

func init() {
	rootCmd.AddCommand(tokenCmd)
	rootCmd.PersistentFlags().StringVar(&tokenName, "name", "", "token名")

}
