package client

import (
	"os"
	"runtime"
	"sync"

	"github.com/spf13/viper"
	"github.com/xavierxcn/apiserver/pkg/log"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var config *rest.Config

// GServiceStore k8s 对象的缓存
var GServiceStore *ServiceStore
var GServiceStoreOnce sync.Once

type ServiceStore struct {
	KubeClient *kubernetes.Clientset
}

// GetServiceStore 获取k8s对象的缓存
func GetServiceStore() *ServiceStore {
	GServiceStoreOnce.Do(func() {
		InitKubeConfig()
	})

	return GServiceStore
}

func InitKubeConfig() {
	kubeConfigPath := ""
	switch runtime.GOOS {
	case "darwin":
		kubeConfigPath = viper.GetString("k3s.kubeconfig.darwin")
	case "windows":
		// kubeConfigPath = viper.GetString("k3s.kubeconfig.windows")
		kubeConfigPath = "D:\\admin.conf"
	case "linux":
		kubeConfigPath = viper.GetString("k3s.kubeconfig.linux")
	default:
		kubeConfigPath = viper.GetString("k3s.kubeconfig.linux")
	}
	if _, err := os.Stat(kubeConfigPath); err == nil {
		config, err = clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	} else {
		log.Errorf("could not get kubeconfig file in: %s", kubeConfigPath)
		return
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	GServiceStore = &ServiceStore{
		KubeClient: clientset,
	}
}

func (ss *ServiceStore) GetKubeClient() *kubernetes.Clientset {
	return ss.KubeClient
}
