package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/xavierxcn/apiserver/internal/serve/boot"
	"github.com/xavierxcn/apiserver/pkg/log"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "api 服务",
	Long:  `api 服务`,
	Run: func(cmd *cobra.Command, args []string) {
		if viper.GetBool("task.enable") {
			boot.TaskBoot()
		}
		boot.ServeBoot()
		log.Info("serve finished")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
