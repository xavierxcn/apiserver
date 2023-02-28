package sd

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xavierxcn/apiserver/pkg/version"
)

const (
	// B b
	B = 1
	// KB kb
	KB = 1024 * B
	// MB mb
	MB = 1024 * KB
	// GB gb
	GB = 1024 * MB
)

// HealthCheck 检查服务是否正常启动的handle
func HealthCheck(c *gin.Context) {
	message := "OK"
	message += fmt.Sprintf("\n" + version.Get() + "\n")

	c.String(http.StatusOK, "\n"+message)
}
