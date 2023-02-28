package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/xavierxcn/apiserver/pkg/version"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "版本",
	Long:  `展示版本信息`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.Get())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
