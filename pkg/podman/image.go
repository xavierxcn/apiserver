package podman

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/pkg/errors"
)

// ImageList returns a list of images
func (c *Client) ImageList() ([]*Image, error) {
	path := "/images/json"

	request, err := http.NewRequest("GET", c.url+path, nil)
	if err != nil {
		return nil, err
	}

	response, err := c.DoHttp(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("list failed: %s", response.Status)
	}

	defer response.Body.Close()

	var images []*Image
	err = json.NewDecoder(response.Body).Decode(&images)
	if err != nil {
		return nil, err
	}

	return images, nil
}

// ImagePull pulls an image from a registry
func (c *Client) ImagePull(image string) error {
	path := "/images/create"
	request, err := http.NewRequest("POST", c.url+path, nil)
	if err != nil {
		return err
	}

	q := request.URL.Query()
	q.Add("fromImage", image)

	request.URL.RawQuery = q.Encode()

	response, err := c.DoHttp(request)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("pull failed: %s", response.Status)
	}

	defer response.Body.Close()
	rdata, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(rdata))

	return nil
}

func cmdLogPath(action string) (string, error) {
	// 创建镜像
	basePath := fmt.Sprintf("/data/upload/%s/%d/", action, time.Now().UnixNano())
	if err := os.MkdirAll(basePath, 0755); err != nil {
		return "", err
	}

	return basePath, nil
}

// ImagePush pushes an image to a registry
func (c *Client) ImagePush(imageTag string) error {
	// 创建镜像
	basePath, err := cmdLogPath("push")
	if err != nil {
		return err
	}

	cmd := exec.Command("podman", "push", imageTag)
	cmdLog, err := os.OpenFile(basePath+"push.log", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return errors.Wrap(err, "create push.log failed")
	}
	cmd.Stdout = cmdLog // 标准输出
	cmd.Stderr = cmdLog // 标准错误
	err = cmd.Start()
	if err != nil {
		return errors.Wrap(err, "start push failed")
	}

	err = cmd.Wait()
	if err != nil {
		return errors.Wrap(err, "push failed")
	}

	return nil
}

// ImageBuild builds an image from a Dockerfile
func (c *Client) ImageBuild(path, imageTag string) error {
	var err error

	// 创建镜像
	basePath, err := cmdLogPath("build")
	if err != nil {
		return err
	}

	cmd := exec.Command("podman", "build", "-t", imageTag, path)
	cmdLog, err := os.OpenFile(basePath+"build.log", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return errors.Wrap(err, "create build.log failed")
	}
	cmd.Stdout = cmdLog // 标准输出
	cmd.Stderr = cmdLog // 标准错误
	err = cmd.Start()
	if err != nil {
		return errors.Wrap(err, "start build failed")
	}

	err = cmd.Wait()
	if err != nil {
		return errors.Wrap(err, "build failed")
	}

	return nil
}
