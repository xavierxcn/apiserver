package auto_update

import (
	"github.com/spf13/viper"
	imagelogic "github.com/xavierxcn/apiserver/internal/task/logic"
	"github.com/xavierxcn/apiserver/internal/task/model/image"
	"github.com/xavierxcn/apiserver/pkg/log"
)

func Run() {
	// 从本地获取，需要更新的镜像
	imagePods, err := image.LoadFromFile(viper.GetString("pod_info.path"))
	if err != nil {
		log.Fatalw("GetAll error", "err", err)
	}

	for _, i := range imagePods {
		err := imagelogic.UpdateVersion(i)
		if err != nil {
			log.Errorw("update version error", "image", i, "err", err)
			continue
		}
	}

}
