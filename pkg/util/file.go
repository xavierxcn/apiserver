// Package util util
package util

import (
	"os"
)

// ABSPath 获取可执行文件当前目录
func ABSPath() (string, error) {
	return os.Getwd()
}
