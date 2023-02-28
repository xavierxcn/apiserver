package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/xavierxcn/apiserver/internal/serve/pkg/config"
	"github.com/xavierxcn/apiserver/pkg/log"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "api",
	Short: "iContainer api 及其相关程序",
	Long:  `iContainer api 及其相关程序`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

var (
	cfgFile string
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "配置文件")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	err := config.Init(cfgFile)
	if err != nil {
		log.Fatal(err.Error())
	}
}
