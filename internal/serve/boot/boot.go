package boot

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	middleware2 "github.com/xavierxcn/apiserver/internal/serve/middleware"
	"github.com/xavierxcn/apiserver/internal/serve/router"
	"github.com/xavierxcn/apiserver/internal/task"
	"github.com/xavierxcn/apiserver/pkg/log"
)

// ServeBoot 启动服务
func ServeBoot() {
	// 设置runmode
	gin.SetMode(viper.GetString("serve.run_mode"))

	// 创建gin实例
	g := gin.New()

	middlewares := []gin.HandlerFunc{
		middleware2.RequestID(),
		middleware2.Logging(),
	}

	router.Load(
		g,
		middlewares...,
	)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatalw("The router has no response, or it might took too long to start up.", "err", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	// 启动https
	cert := viper.GetString("tls.cert")
	key := viper.GetString("tls.key")
	if cert != "" && key != "" {
		go func() {
			log.Infof("Start to listening the incoming requests on https address: %s", viper.GetString("tls.addr"))
			log.Infof(http.ListenAndServeTLS(viper.GetString("tls.addr"), cert, key, g).Error())
		}()
	}
	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	if err := g.Run(viper.GetString("serve.addr")); err != nil {
		log.Fatalw("An error has occurred while starting the gin instance.", "err", err)
	}
}

// TaskBoot 启动任务
func TaskBoot() {
	task.Run()
}

// 检查是否正常启动
func pingServer() error {
	for i := 0; i < viper.GetInt("serve.health.max_ping_count"); i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get(viper.GetString("serve.health.url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Infof("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router")
}
