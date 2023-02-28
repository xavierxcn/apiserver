/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/xavierxcn/apiserver/internal/serve/model"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "数据库迁移",
	Long:  `数据库迁移，使用gorm的migrate`,
	Run: func(cmd *cobra.Command, args []string) {
		model.AutoMigrate()
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
