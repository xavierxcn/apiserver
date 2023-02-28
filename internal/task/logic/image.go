package image

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"github.com/xavierxcn/apiserver/internal/task/model/image"
	"github.com/xavierxcn/apiserver/pkg/deployment"
	"github.com/xavierxcn/apiserver/pkg/log"
)

// UpdateVersion update image version
func UpdateVersion(i *image.PodInfo) error {
	var err error
	// todo 从云端获取最新的版本
	imageLatest := &imageVersion{
		ImageName: i.ImageName,
	}
	err = imageLatest.getLatestVersion()
	if err != nil {
		return err
	}

	log.Infow("get registry version", "image", i.ImageName)
	log.Infow("get registry version successfully",
		"image", imageLatest.ImageName,
		"version", imageLatest.Version)

	// 获取k3s版本
	log.Infow("get k3s version", "namespace", i.Namespace)
	registry, k3sImage, k3sVersion, err := getK3sVersion(i.Namespace, i.PodName)
	if err != nil {
		return errors.Wrapf(err, "get k3s version error")
	}
	log.Infow("get k3s version successfully",
		"namespace", i.Namespace,
		"registry", registry, "image", k3sImage, "version", k3sVersion)

	// 比较版本
	if k3sVersion == imageLatest.Version {
		log.Infow("version is same", "image", imageLatest.ImageName, "version", imageLatest.Version)
		return nil
	}

	newK3sImageVersion := fmt.Sprintf("%s:%s", imageLatest.ImageName, imageLatest.Version)

	log.Infow("获取到镜像更新", "image", i, "old", k3sVersion, "new", imageLatest.Version)

	// 更新k3s版本
	log.Infow("更新k3s版本", "image", i, "old", k3sVersion, "new", imageLatest.Version)
	err = deployment.UpdateVersion(i.Namespace, i.PodName, newK3sImageVersion)
	if err != nil {
		return errors.Wrapf(err, "update k3s version error")
	}
	log.Infow("更新k3s版本成功", "image", i, "old", k3sVersion, "new", imageLatest.Version)

	log.Infow("镜像更新成功", "image", i, "version", imageLatest.Version, "new_version", newK3sImageVersion)
	return nil
}

type TagList struct {
	Tags []string `json:"tags"`
	Name string   `json:"name"`
}

// imageVersion 镜像版本
type imageVersion struct {
	ImageName string `json:"imageName"` // 镜像名称
	Version   string `json:"version"`   // 版本
}

type imageVersionRsp struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    *imageVersion `json:"data"`
}

func (iv *imageVersion) getLatestVersion() error {
	path := "/v1/apiserver/image/image/version/latest"
	registryUrl := viper.GetString("task.server_url")

	reqData, err := json.Marshal(iv)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", registryUrl+path, bytes.NewBuffer(reqData))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", "Bearer "+viper.GetString("task.token"))
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	rspData := &imageVersionRsp{}
	err = json.NewDecoder(response.Body).Decode(rspData)
	if err != nil {
		return err
	}

	if rspData.Code != 0 {
		return errors.New(rspData.Message)
	}

	iv.Version = rspData.Data.Version

	return nil
}

func getK3sVersion(namespace, name string) (registry, imageName, imageVersion string, err error) {
	k3sImageVersion, err := deployment.GetVersion(namespace, name)
	if err != nil {
		return
	}

	registry, imageName, imageVersion, err = parseK3sImage(k3sImageVersion)
	if err != nil {
		return
	}

	return
}

func parseK3sImage(image string) (registry string, name string, version string, err error) {
	s := strings.SplitN(image, "/", 2)
	if len(s) < 2 {
		err = fmt.Errorf("invalid image: %s", image)
		return
	}

	s2 := strings.SplitN(s[1], ":", 2)
	if len(s2) < 2 {
		err = fmt.Errorf("invalid image: %s", image)
		return
	}

	return s[0], s2[0], s2[1], nil
}
