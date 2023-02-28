package router

import (
	apiv1 "github.com/xavierxcn/apiserver/api/v1"
	"github.com/xavierxcn/apiserver/internal/serve/handler/hello"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xavierxcn/apiserver/internal/serve/handler/sd"
	"github.com/xavierxcn/apiserver/internal/serve/middleware"
)

// Load load
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})

	// 健康检查
	svcd := g.Group("/sd")
	{
		svcd.GET("/health", sd.HealthCheck)
	}

	v1 := g.Group("/v1")
	{
		apiv1.RegisterHelloServiceHTTPServer(v1, hello.Service{})
	}

	return g

}
