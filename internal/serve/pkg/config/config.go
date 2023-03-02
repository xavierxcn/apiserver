package config

import (
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/xavierxcn/apiserver/pkg/log"
)

// Config 配置
type Config struct {
	Name string
}

var (
	// 默认配置
	defaultConfig = map[string]interface{}{
		// serve
		"serve.run_mode":              "debug",
		"serve.addr":                  "0.0.0.0:8765",
		"serve.health.url":            "http://127.0.0.1:8765",
		"serve.health.max_ping_count": 10,
		"serve.iregistry.enable":      true,
		"serve.jwt_secret":            "",

		// log
		"log.level":  "debug",
		"log.format": "console",
		"log.output": []string{"stdout"},

		// podman
		"podman.sock":    "/var/run/podman/podman.sock",
		"podman.timeout": 10,

		// registry
		"registry.url":      "http://172.16.8.10:5000",
		"registry.username": "",
		"registry.password": "",
	}
)

// Init 初始化
func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	// 初始化配置文件
	if err := c.initConfig(); err != nil {
		return errors.Wrap(err, "init config failed")
	}

	// 初始化日志
	c.initLog()

	// 监控配置文件的变化并热加载
	c.watchConfig()

	return nil
}

// 初始化配置
func (c *Config) initConfig() error {
	if c.Name != "" {
		// 指定了配置文件路径，则解析指定的配置文件
		viper.SetConfigFile(c.Name)
		viper.SetConfigType("yaml")

		// 解析所有配置信息
		if err := viper.ReadInConfig(); err != nil {
			return errors.Wrap(err, "read config failed")
		}
	}

	// 需要读取的环境变量的前缀
	viper.SetEnvPrefix("APISERVER")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	// 设置读取环境变量中的配置
	viper.AutomaticEnv()

	for k, v := range defaultConfig {
		if !viper.IsSet(k) {
			viper.Set(k, v)
		}
	}

	return nil
}

// 初始化log
func (c *Config) initLog() {
	// logger配置
	opts := &log.Options{
		Level:            viper.GetString("log.level"),
		Format:           viper.GetString("log.format"),
		EnableColor:      false,
		DisableCaller:    false,
		OutputPaths:      viper.GetStringSlice("log.output"),
		ErrorOutputPaths: []string{"stderr"},
	}
	// 初始化全局logger
	log.Init(opts)
}

// 监控配置文件的变动，实时更新
func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Warnf("Config file changed: %s", e.Name)
	})
}
