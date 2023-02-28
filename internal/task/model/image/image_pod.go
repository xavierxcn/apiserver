package image

import (
	"encoding/json"
	"os"
)

// PodInfo 镜像版本
type PodInfo struct {
	ImageName string `json:"imageName"` // 镜像名称
	PodName   string `json:"podName"`   // pod名称
	Namespace string `json:"namespace"` // 命名空间
}

// LoadFromFile 加载镜像版本信息
func LoadFromFile(path string) ([]*PodInfo, error) {
	// 从json文件中读取
	if path == "" {
		path = "./pod_info.json"
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	var podInfos []*PodInfo
	err = json.NewDecoder(file).Decode(&podInfos)
	if err != nil {
		return nil, err
	}
	return podInfos, nil
}
