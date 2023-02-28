package deployment

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/xavierxcn/apiserver/pkg/client"
	"github.com/xavierxcn/apiserver/pkg/log"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func Create(namespace string, deploy *v1.Deployment) error {
	cw := client.GetServiceStore().GetKubeClient()
	_, err := cw.AppsV1().Deployments(namespace).Create(context.TODO(), deploy, metav1.CreateOptions{})
	if err != nil {
		log.Errorf("create deploy error %s", err.Error())
		return err
	}
	return nil
}

func Update(namespace string, deploy *v1.Deployment) error {
	cw := client.GetServiceStore().GetKubeClient()
	_, err := cw.AppsV1().Deployments(namespace).Update(context.TODO(), deploy, metav1.UpdateOptions{})
	if err != nil {
		log.Errorf("create deploy error %s", err.Error())
		return err
	}
	return nil
}

func Delete(namespace, name string) error {
	cw := client.GetServiceStore().GetKubeClient()
	err := cw.AppsV1().Deployments(namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		log.Errorf("create deploy error %s", err.Error())
		return err
	}
	return nil
}

// UpdateVersion 通过deployment的namespace，name 和 version 更新deployment
func UpdateVersion(namespace, name, version string) error {
	if namespace == "" || name == "" || version == "" {
		return errors.New("namespace,name,version is null")
	}
	cw := client.GetServiceStore().GetKubeClient()
	dp, err := cw.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return errors.New(fmt.Sprintf("get deployment %s error: %s", name, err.Error()))
	}
	if len(dp.Spec.Template.Spec.Containers) == 0 {
		return errors.New("deployment is nil")
	}
	dp.Spec.Template.Spec.Containers[0].Image = version
	_, err2 := cw.AppsV1().Deployments(namespace).Update(context.TODO(), dp, metav1.UpdateOptions{})
	if err2 != nil {
		return errors.New(fmt.Sprintf("update deployment %s error: %s", name, err2.Error()))
	}
	return nil
}

// GetVersion 通过deployment的namespace和name获取镜像地址
func GetVersion(namespace, name string) (string, error) {
	if namespace == "" || name == "" {
		return "", errors.New("namespace,name is null")
	}
	cw := client.GetServiceStore().GetKubeClient()
	dp, err := cw.AppsV1().Deployments(namespace).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return "", errors.New(fmt.Sprintf("get deployment %s error: %s", name, err.Error()))
	}
	if len(dp.Spec.Template.Spec.Containers) == 0 {
		return "", errors.New("deployment is nil")
	}
	return dp.Spec.Template.Spec.Containers[0].Image, nil
}
